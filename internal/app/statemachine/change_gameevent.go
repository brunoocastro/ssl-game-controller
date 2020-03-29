package statemachine

import (
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/config"
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/statemachine/ballplace"
	"github.com/golang/protobuf/ptypes"
	"log"
)

func (s *StateMachine) AddGameEvent(newState *state.State, change *AddGameEvent) (changes []*Change) {

	gameEvent := change.GameEvent
	byTeam := gameEvent.ByTeam()

	if gameEvent.Type == nil {
		log.Printf("Can not process a game event without a type: %v", gameEvent)
		return
	}

	// convert aimless kick if necessary
	if newState.Division.Div() == config.DivA && *gameEvent.Type == state.GameEvent_AIMLESS_KICK {
		log.Println("Convert aimless kick to ball left field event, because we are in DivA")
		gameEvent = s.convertAimlessKick(change.GameEvent)
	}

	// remember game event
	newState.GameEvents = append(newState.GameEvents, gameEvent)

	// determine next command
	newState.NextCommand = s.nextCommandForEvent(newState, gameEvent)

	// Increment foul counter
	if incrementsFoulCounter(*gameEvent.Type) {
		for _, team := range state.BothTeams() {
			if byTeam == state.Team_UNKNOWN || byTeam == team {
				log.Printf("Team %v got a foul for %v", byTeam, gameEvent)
				newState.TeamInfo(team).AddFoul(gameEvent)
				if len(newState.TeamInfo(team).Fouls)%3 == 0 {
					changes = append(changes, s.multipleFoulsChange(team))
				}
			}
		}
	}

	// Add yellow card
	if addsYellowCard(*gameEvent.Type) && byTeam.Known() {
		log.Printf("Team %v got a yellow card", byTeam)
		changes = append(changes, &Change{
			Change: &Change_AddYellowCard{
				AddYellowCard: &AddYellowCard{
					ForTeam:           &byTeam,
					CausedByGameEvent: gameEvent,
				},
			},
		})
	}

	// Add red card
	if addsRedCard(*gameEvent.Type) && byTeam.Known() {
		log.Printf("Team %v got a red card", byTeam)
		changes = append(changes, &Change{
			Change: &Change_AddRedCard{
				AddRedCard: &AddRedCard{
					ForTeam:           &byTeam,
					CausedByGameEvent: gameEvent,
				},
			},
		})
	}

	// goal
	if *gameEvent.Type == state.GameEvent_GOAL && byTeam.Known() {
		*newState.TeamInfo(byTeam).Goals++
	}

	// possible goal
	if *gameEvent.Type == state.GameEvent_POSSIBLE_GOAL {
		log.Printf("Halt the game, because team %v might have scored a goal", byTeam)
		// halt the game to let the human referee decide if this was a valid goal
		changes = append(changes, s.newCommandChange(state.NewCommandNeutral(state.Command_HALT)))
	}

	// bot substitution
	if *gameEvent.Type == state.GameEvent_BOT_SUBSTITUTION {
		log.Printf("Halt the game, because team %v requested robot substitution", byTeam)
		// reset robot substitution flags
		for _, team := range state.BothTeams() {
			*newState.TeamInfo(team).BotSubstitutionIntent = false
		}
		// halt the game to allow teams to substitute robots
		changes = append(changes, s.newCommandChange(state.NewCommandNeutral(state.Command_HALT)))
	}

	// ball placement interference
	if *gameEvent.Type == state.GameEvent_BOT_INTERFERED_PLACEMENT {
		log.Printf("Reset current action time for ball placement interference by %v", byTeam)
		curDuration, _ := ptypes.Duration(newState.CurrentActionTimeRemaining)
		newState.CurrentActionTimeRemaining = ptypes.DurationProto(curDuration + s.gameConfig.BallPlacementTimeTopUp)
	}

	// ball placement position
	placementPosDeterminer := ballplace.BallPlacementPosDeterminer{
		Event:               gameEvent,
		Geometry:            s.geometry,
		CurrentPlacementPos: newState.PlacementPos,
		OnPositiveHalf: map[state.Team]bool{
			state.Team_BLUE:   *newState.TeamInfo(state.Team_BLUE).OnPositiveHalf,
			state.Team_YELLOW: *newState.TeamInfo(state.Team_YELLOW).OnPositiveHalf,
		},
	}
	newState.PlacementPos = placementPosDeterminer.Location()

	// ball placement failed
	if *gameEvent.Type == state.GameEvent_PLACEMENT_FAILED && byTeam.Known() {
		*newState.TeamInfo(byTeam).BallPlacementFailures++
		*newState.TeamInfo(byTeam).BallPlacementFailuresReached = *newState.TeamInfo(byTeam).BallPlacementFailures >= s.gameConfig.MultiplePlacementFailures
		if s.allTeamsFailedPlacement(newState) {
			log.Printf("Placement failed for all teams. The human ref must place the ball.")
			changes = append(changes, s.newCommandChange(state.NewCommandNeutral(state.Command_HALT)))
		} else {
			log.Printf("Placement failed for team %v. Team %v is awarded a free kick and places the ball.", byTeam, byTeam.Opposite())
			newState.NextCommand = state.NewCommand(state.Command_DIRECT, byTeam.Opposite())
			changes = append(changes, s.newCommandChange(state.NewCommand(state.Command_BALL_PLACEMENT, byTeam.Opposite())))
		}
	}

	// ball placement succeeded
	if *gameEvent.Type == state.GameEvent_PLACEMENT_SUCCEEDED &&
		byTeam.Known() &&
		*newState.TeamInfo(byTeam).BallPlacementFailures > 0 {
		*newState.TeamInfo(byTeam).BallPlacementFailures--
		if byTeam == *newState.NextCommand.ForTeam {
			log.Printf("Placement succeeded by team %v, which is also in favor. Can continue.", byTeam)
			changes = append(changes, &Change{
				Change: &Change_Continue{
					Continue: &Continue{},
				},
			})
		}
	}

	// defender too close to kick point
	if *gameEvent.Type == state.GameEvent_DEFENDER_TOO_CLOSE_TO_KICK_POINT {
		log.Printf("Reset current action time because defender of team %v was too close to kick point", byTeam)
		newState.CurrentActionTimeRemaining = ptypes.DurationProto(s.gameConfig.GeneralTime)
	}

	// stop the game if needed
	if *newState.Command.Type != state.Command_STOP &&
		stopsTheGame(*gameEvent.Type) {
		log.Printf("Stopping the game for event %v", *gameEvent.Type)
		changes = append(changes, s.newCommandChange(state.NewCommandNeutral(state.Command_STOP)))
	}

	return
}

// multipleFoulsChange creates a multiple fouls event change
func (s *StateMachine) multipleFoulsChange(byTeam state.Team) *Change {
	eventType := state.GameEvent_MULTIPLE_FOULS
	return &Change{
		Change: &Change_AddGameEvent{
			AddGameEvent: &AddGameEvent{
				GameEvent: &state.GameEvent{
					Type: &eventType,
					Event: &state.GameEvent_MultipleFouls_{
						MultipleFouls: &state.GameEvent_MultipleFouls{
							ByTeam: &byTeam,
						},
					},
				},
			},
		},
	}
}

// convertAimlessKick converts the aimless kick event into a ball left field via goal line event
// because aimless kick only applies to DivB
func (s *StateMachine) convertAimlessKick(gameEvent *state.GameEvent) *state.GameEvent {
	eventType := state.GameEvent_BALL_LEFT_FIELD_GOAL_LINE
	return &state.GameEvent{
		Type:   &eventType,
		Origin: gameEvent.Origin,
		Event: &state.GameEvent_BallLeftFieldGoalLine{
			BallLeftFieldGoalLine: &state.GameEvent_BallLeftField{
				ByTeam:   gameEvent.GetAimlessKick().ByTeam,
				ByBot:    gameEvent.GetAimlessKick().ByBot,
				Location: gameEvent.GetAimlessKick().Location,
			},
		},
	}
}

// nextCommandForEvent determines the next command for the given event or returns the currently set one
func (s *StateMachine) nextCommandForEvent(newState *state.State, gameEvent *state.GameEvent) (command *state.Command) {
	if *newState.Command.Type == state.Command_PENALTY || *newState.Command.Type == state.Command_KICKOFF {
		return state.NewCommand(state.Command_NORMAL_START, state.Team_UNKNOWN)
	}

	switch *gameEvent.Type {
	case state.GameEvent_BALL_LEFT_FIELD_GOAL_LINE,
		state.GameEvent_BALL_LEFT_FIELD_TOUCH_LINE,
		state.GameEvent_AIMLESS_KICK,
		state.GameEvent_ATTACKER_TOO_CLOSE_TO_DEFENSE_AREA,
		state.GameEvent_BOT_PUSHED_BOT,
		state.GameEvent_BOT_HELD_BALL_DELIBERATELY,
		state.GameEvent_BOT_TIPPED_OVER,
		state.GameEvent_KEEPER_HELD_BALL,
		state.GameEvent_BOUNDARY_CROSSING,
		state.GameEvent_BOT_DRIBBLED_BALL_TOO_FAR,
		state.GameEvent_ATTACKER_DOUBLE_TOUCHED_BALL,
		state.GameEvent_POSSIBLE_GOAL:
		return state.NewCommand(state.Command_DIRECT, gameEvent.ByTeam().Opposite())
	case state.GameEvent_DEFENDER_IN_DEFENSE_AREA:
		return state.NewCommand(state.Command_PENALTY, gameEvent.ByTeam().Opposite())
	case state.GameEvent_NO_PROGRESS_IN_GAME:
		return state.NewCommand(state.Command_FORCE_START, state.Team_UNKNOWN)
	default:
		return newState.NextCommand
	}
}

// incrementsFoulCounter checks if the game event increments the foul counter
func incrementsFoulCounter(gameEvent state.GameEvent_Type) bool {
	switch gameEvent {
	case
		state.GameEvent_AIMLESS_KICK,
		state.GameEvent_KEEPER_HELD_BALL,
		state.GameEvent_ATTACKER_TOUCHED_BALL_IN_DEFENSE_AREA,
		state.GameEvent_BOT_DRIBBLED_BALL_TOO_FAR,
		state.GameEvent_BOT_KICKED_BALL_TOO_FAST,
		state.GameEvent_ATTACKER_TOO_CLOSE_TO_DEFENSE_AREA,
		state.GameEvent_BOT_INTERFERED_PLACEMENT,
		state.GameEvent_BOT_CRASH_DRAWN,
		state.GameEvent_BOT_CRASH_UNIQUE,
		state.GameEvent_BOT_PUSHED_BOT,
		state.GameEvent_BOT_HELD_BALL_DELIBERATELY,
		state.GameEvent_BOT_TIPPED_OVER,
		state.GameEvent_BOT_TOO_FAST_IN_STOP,
		state.GameEvent_DEFENDER_TOO_CLOSE_TO_KICK_POINT,
		state.GameEvent_BOUNDARY_CROSSING:
		return true
	}
	return false
}

// addsYellowCard checks if the game event adds a yellow card
func addsYellowCard(gameEvent state.GameEvent_Type) bool {
	switch gameEvent {
	case
		state.GameEvent_MULTIPLE_FOULS,
		state.GameEvent_UNSPORTING_BEHAVIOR_MINOR:
		return true
	}
	return false
}

// addsYellowCard checks if the game event adds a yellow card
func addsRedCard(gameEvent state.GameEvent_Type) bool {
	switch gameEvent {
	case
		state.GameEvent_DEFENDER_IN_DEFENSE_AREA,
		state.GameEvent_UNSPORTING_BEHAVIOR_MAJOR:
		return true
	}
	return false
}

// addsYellowCard checks if the game event adds a yellow card
func stopsTheGame(gameEvent state.GameEvent_Type) bool {
	switch gameEvent {
	case
		state.GameEvent_ATTACKER_TOO_CLOSE_TO_DEFENSE_AREA,
		state.GameEvent_BOT_PUSHED_BOT,
		state.GameEvent_BOT_HELD_BALL_DELIBERATELY,
		state.GameEvent_BOT_TIPPED_OVER,
		state.GameEvent_DEFENDER_IN_DEFENSE_AREA,
		state.GameEvent_BOUNDARY_CROSSING,
		state.GameEvent_KEEPER_HELD_BALL,
		state.GameEvent_BOT_DRIBBLED_BALL_TOO_FAR,
		state.GameEvent_PLACEMENT_SUCCEEDED:
		return true
	}
	return false
}

// allTeamsFailedPlacement returns true if all teams failed placing the ball
// It takes into account, how many teams are able to place the ball and how many failures happened
func (s *StateMachine) allTeamsFailedPlacement(newState *state.State) bool {
	possibleFailures := 0
	for _, team := range state.BothTeams() {
		if newState.TeamInfo(team).BallPlacementAllowed() {
			possibleFailures++
		}
	}

	failures := 0
	for _, e := range newState.GameEvents {
		if *e.Type == state.GameEvent_PLACEMENT_FAILED {
			failures++
			if failures >= possibleFailures {
				return true
			}
		}
	}
	return false
}
