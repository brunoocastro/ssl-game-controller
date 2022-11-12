// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: ssl_gc_engine.proto

package engine

import (
	geom "github.com/RoboCup-SSL/ssl-game-controller/internal/app/geom"
	state "github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// possible advantage choices
type TeamAdvantageChoice_AdvantageChoice int32

const (
	// stop the game
	TeamAdvantageChoice_STOP TeamAdvantageChoice_AdvantageChoice = 0
	// keep the match running
	TeamAdvantageChoice_CONTINUE TeamAdvantageChoice_AdvantageChoice = 1
)

// Enum value maps for TeamAdvantageChoice_AdvantageChoice.
var (
	TeamAdvantageChoice_AdvantageChoice_name = map[int32]string{
		0: "STOP",
		1: "CONTINUE",
	}
	TeamAdvantageChoice_AdvantageChoice_value = map[string]int32{
		"STOP":     0,
		"CONTINUE": 1,
	}
)

func (x TeamAdvantageChoice_AdvantageChoice) Enum() *TeamAdvantageChoice_AdvantageChoice {
	p := new(TeamAdvantageChoice_AdvantageChoice)
	*p = x
	return p
}

func (x TeamAdvantageChoice_AdvantageChoice) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TeamAdvantageChoice_AdvantageChoice) Descriptor() protoreflect.EnumDescriptor {
	return file_ssl_gc_engine_proto_enumTypes[0].Descriptor()
}

func (TeamAdvantageChoice_AdvantageChoice) Type() protoreflect.EnumType {
	return &file_ssl_gc_engine_proto_enumTypes[0]
}

func (x TeamAdvantageChoice_AdvantageChoice) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TeamAdvantageChoice_AdvantageChoice) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TeamAdvantageChoice_AdvantageChoice(num)
	return nil
}

// Deprecated: Use TeamAdvantageChoice_AdvantageChoice.Descriptor instead.
func (TeamAdvantageChoice_AdvantageChoice) EnumDescriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{2, 0}
}

// The GC state contains settings and state independent of the match state
type GcState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state of each team
	TeamState map[string]*GcStateTeam `protobuf:"bytes,1,rep,name=team_state,json=teamState" json:"team_state,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// the states of the auto referees
	AutoRefState map[string]*GcStateAutoRef `protobuf:"bytes,2,rep,name=auto_ref_state,json=autoRefState" json:"auto_ref_state,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// the states of the attached trackers
	TrackerState map[string]*GcStateTracker `protobuf:"bytes,3,rep,name=tracker_state,json=trackerState" json:"tracker_state,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// the state of the currently selected tracker
	TrackerStateGc *GcStateTracker `protobuf:"bytes,4,opt,name=tracker_state_gc,json=trackerStateGc" json:"tracker_state_gc,omitempty"`
	// can the match be continued right now?
	ReadyToContinue *bool `protobuf:"varint,5,opt,name=ready_to_continue,json=readyToContinue" json:"ready_to_continue,omitempty"`
	// list of issues that hinders the game from continuing
	ContinuationIssues []string `protobuf:"bytes,6,rep,name=continuation_issues,json=continuationIssues" json:"continuation_issues,omitempty"`
}

func (x *GcState) Reset() {
	*x = GcState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcState) ProtoMessage() {}

func (x *GcState) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcState.ProtoReflect.Descriptor instead.
func (*GcState) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{0}
}

func (x *GcState) GetTeamState() map[string]*GcStateTeam {
	if x != nil {
		return x.TeamState
	}
	return nil
}

func (x *GcState) GetAutoRefState() map[string]*GcStateAutoRef {
	if x != nil {
		return x.AutoRefState
	}
	return nil
}

func (x *GcState) GetTrackerState() map[string]*GcStateTracker {
	if x != nil {
		return x.TrackerState
	}
	return nil
}

func (x *GcState) GetTrackerStateGc() *GcStateTracker {
	if x != nil {
		return x.TrackerStateGc
	}
	return nil
}

func (x *GcState) GetReadyToContinue() bool {
	if x != nil && x.ReadyToContinue != nil {
		return *x.ReadyToContinue
	}
	return false
}

func (x *GcState) GetContinuationIssues() []string {
	if x != nil {
		return x.ContinuationIssues
	}
	return nil
}

// The GC state for a single team
type GcStateTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// true: The team is connected
	Connected *bool `protobuf:"varint,1,opt,name=connected" json:"connected,omitempty"`
	// true: The remote control for the team is connected
	RemoteControlConnected *bool `protobuf:"varint,3,opt,name=remote_control_connected,json=remoteControlConnected" json:"remote_control_connected,omitempty"`
	// the advantage choice of the team
	AdvantageChoice *TeamAdvantageChoice `protobuf:"bytes,5,opt,name=advantage_choice,json=advantageChoice" json:"advantage_choice,omitempty"`
}

func (x *GcStateTeam) Reset() {
	*x = GcStateTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcStateTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcStateTeam) ProtoMessage() {}

func (x *GcStateTeam) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcStateTeam.ProtoReflect.Descriptor instead.
func (*GcStateTeam) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{1}
}

func (x *GcStateTeam) GetConnected() bool {
	if x != nil && x.Connected != nil {
		return *x.Connected
	}
	return false
}

func (x *GcStateTeam) GetRemoteControlConnected() bool {
	if x != nil && x.RemoteControlConnected != nil {
		return *x.RemoteControlConnected
	}
	return false
}

func (x *GcStateTeam) GetAdvantageChoice() *TeamAdvantageChoice {
	if x != nil {
		return x.AdvantageChoice
	}
	return nil
}

// The choice from a team regarding the advantage rule
type TeamAdvantageChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the choice of the team
	Choice *TeamAdvantageChoice_AdvantageChoice `protobuf:"varint,1,opt,name=choice,enum=TeamAdvantageChoice_AdvantageChoice" json:"choice,omitempty"`
}

func (x *TeamAdvantageChoice) Reset() {
	*x = TeamAdvantageChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamAdvantageChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamAdvantageChoice) ProtoMessage() {}

func (x *TeamAdvantageChoice) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamAdvantageChoice.ProtoReflect.Descriptor instead.
func (*TeamAdvantageChoice) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{2}
}

func (x *TeamAdvantageChoice) GetChoice() TeamAdvantageChoice_AdvantageChoice {
	if x != nil && x.Choice != nil {
		return *x.Choice
	}
	return TeamAdvantageChoice_STOP
}

// The GC state of an auto referee
type GcStateAutoRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GcStateAutoRef) Reset() {
	*x = GcStateAutoRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcStateAutoRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcStateAutoRef) ProtoMessage() {}

func (x *GcStateAutoRef) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcStateAutoRef.ProtoReflect.Descriptor instead.
func (*GcStateAutoRef) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{3}
}

// GC state of a tracker
type GcStateTracker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the source
	SourceName *string `protobuf:"bytes,1,opt,name=source_name,json=sourceName" json:"source_name,omitempty"`
	// UUID of the source
	Uuid *string `protobuf:"bytes,4,opt,name=uuid" json:"uuid,omitempty"`
	// Current ball
	Ball *Ball `protobuf:"bytes,2,opt,name=ball" json:"ball,omitempty"`
	// Current robots
	Robots []*Robot `protobuf:"bytes,3,rep,name=robots" json:"robots,omitempty"`
}

func (x *GcStateTracker) Reset() {
	*x = GcStateTracker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcStateTracker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcStateTracker) ProtoMessage() {}

func (x *GcStateTracker) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcStateTracker.ProtoReflect.Descriptor instead.
func (*GcStateTracker) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{4}
}

func (x *GcStateTracker) GetSourceName() string {
	if x != nil && x.SourceName != nil {
		return *x.SourceName
	}
	return ""
}

func (x *GcStateTracker) GetUuid() string {
	if x != nil && x.Uuid != nil {
		return *x.Uuid
	}
	return ""
}

func (x *GcStateTracker) GetBall() *Ball {
	if x != nil {
		return x.Ball
	}
	return nil
}

func (x *GcStateTracker) GetRobots() []*Robot {
	if x != nil {
		return x.Robots
	}
	return nil
}

// The ball state
type Ball struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ball position [m]
	Pos *geom.Vector3 `protobuf:"bytes,1,opt,name=pos" json:"pos,omitempty"`
	// ball velocity [m/s]
	Vel *geom.Vector3 `protobuf:"bytes,2,opt,name=vel" json:"vel,omitempty"`
}

func (x *Ball) Reset() {
	*x = Ball{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ball) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ball) ProtoMessage() {}

func (x *Ball) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ball.ProtoReflect.Descriptor instead.
func (*Ball) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{5}
}

func (x *Ball) GetPos() *geom.Vector3 {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *Ball) GetVel() *geom.Vector3 {
	if x != nil {
		return x.Vel
	}
	return nil
}

// The robot state
type Robot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// robot id and team
	Id *state.RobotId `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// robot position [m]
	Pos *geom.Vector2 `protobuf:"bytes,2,opt,name=pos" json:"pos,omitempty"`
}

func (x *Robot) Reset() {
	*x = Robot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Robot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Robot) ProtoMessage() {}

func (x *Robot) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Robot.ProtoReflect.Descriptor instead.
func (*Robot) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{6}
}

func (x *Robot) GetId() *state.RobotId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Robot) GetPos() *geom.Vector2 {
	if x != nil {
		return x.Pos
	}
	return nil
}

var File_ssl_gc_engine_proto protoreflect.FileDescriptor

var file_ssl_gc_engine_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x67, 0x65,
	0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x73,
	0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcc, 0x04, 0x0a, 0x07, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a,
	0x0a, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x74, 0x65, 0x61, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65,
	0x66, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x52,
	0x65, 0x66, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x67, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x47, 0x63, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x74, 0x6f, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x12,
	0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f,
	0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73,
	0x1a, 0x4a, 0x0a, 0x0e, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50, 0x0a, 0x11,
	0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f,
	0x52, 0x65, 0x66, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50,
	0x0a, 0x11, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xa6, 0x01, 0x0a, 0x0b, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x38,
	0x0a, 0x18, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x16, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x10, 0x61, 0x64, 0x76, 0x61,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x0f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x7e, 0x0a, 0x13, 0x54, 0x65, 0x61,
	0x6d, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x24, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x29,
	0x0a, 0x0f, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45, 0x10, 0x01, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66, 0x22, 0x80, 0x01, 0x0a, 0x0e,
	0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x62, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x52, 0x04, 0x62, 0x61, 0x6c, 0x6c, 0x12, 0x1e,
	0x0a, 0x06, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x52, 0x06, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x22, 0x3e,
	0x0a, 0x04, 0x42, 0x61, 0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x03, 0x70,
	0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x03, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x03, 0x76, 0x65, 0x6c, 0x22, 0x3d,
	0x0a, 0x05, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x42, 0x40, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f,
	0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d, 0x67, 0x61, 0x6d, 0x65,
	0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
}

var (
	file_ssl_gc_engine_proto_rawDescOnce sync.Once
	file_ssl_gc_engine_proto_rawDescData = file_ssl_gc_engine_proto_rawDesc
)

func file_ssl_gc_engine_proto_rawDescGZIP() []byte {
	file_ssl_gc_engine_proto_rawDescOnce.Do(func() {
		file_ssl_gc_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_gc_engine_proto_rawDescData)
	})
	return file_ssl_gc_engine_proto_rawDescData
}

var file_ssl_gc_engine_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ssl_gc_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_ssl_gc_engine_proto_goTypes = []interface{}{
	(TeamAdvantageChoice_AdvantageChoice)(0), // 0: TeamAdvantageChoice.AdvantageChoice
	(*GcState)(nil),                          // 1: GcState
	(*GcStateTeam)(nil),                      // 2: GcStateTeam
	(*TeamAdvantageChoice)(nil),              // 3: TeamAdvantageChoice
	(*GcStateAutoRef)(nil),                   // 4: GcStateAutoRef
	(*GcStateTracker)(nil),                   // 5: GcStateTracker
	(*Ball)(nil),                             // 6: Ball
	(*Robot)(nil),                            // 7: Robot
	nil,                                      // 8: GcState.TeamStateEntry
	nil,                                      // 9: GcState.AutoRefStateEntry
	nil,                                      // 10: GcState.TrackerStateEntry
	(*geom.Vector3)(nil),                     // 11: Vector3
	(*state.RobotId)(nil),                    // 12: RobotId
	(*geom.Vector2)(nil),                     // 13: Vector2
}
var file_ssl_gc_engine_proto_depIdxs = []int32{
	8,  // 0: GcState.team_state:type_name -> GcState.TeamStateEntry
	9,  // 1: GcState.auto_ref_state:type_name -> GcState.AutoRefStateEntry
	10, // 2: GcState.tracker_state:type_name -> GcState.TrackerStateEntry
	5,  // 3: GcState.tracker_state_gc:type_name -> GcStateTracker
	3,  // 4: GcStateTeam.advantage_choice:type_name -> TeamAdvantageChoice
	0,  // 5: TeamAdvantageChoice.choice:type_name -> TeamAdvantageChoice.AdvantageChoice
	6,  // 6: GcStateTracker.ball:type_name -> Ball
	7,  // 7: GcStateTracker.robots:type_name -> Robot
	11, // 8: Ball.pos:type_name -> Vector3
	11, // 9: Ball.vel:type_name -> Vector3
	12, // 10: Robot.id:type_name -> RobotId
	13, // 11: Robot.pos:type_name -> Vector2
	2,  // 12: GcState.TeamStateEntry.value:type_name -> GcStateTeam
	4,  // 13: GcState.AutoRefStateEntry.value:type_name -> GcStateAutoRef
	5,  // 14: GcState.TrackerStateEntry.value:type_name -> GcStateTracker
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_ssl_gc_engine_proto_init() }
func file_ssl_gc_engine_proto_init() {
	if File_ssl_gc_engine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ssl_gc_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ssl_gc_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcStateTeam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ssl_gc_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamAdvantageChoice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ssl_gc_engine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcStateAutoRef); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ssl_gc_engine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcStateTracker); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ssl_gc_engine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ball); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ssl_gc_engine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Robot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ssl_gc_engine_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_gc_engine_proto_goTypes,
		DependencyIndexes: file_ssl_gc_engine_proto_depIdxs,
		EnumInfos:         file_ssl_gc_engine_proto_enumTypes,
		MessageInfos:      file_ssl_gc_engine_proto_msgTypes,
	}.Build()
	File_ssl_gc_engine_proto = out.File
	file_ssl_gc_engine_proto_rawDesc = nil
	file_ssl_gc_engine_proto_goTypes = nil
	file_ssl_gc_engine_proto_depIdxs = nil
}
