package rcon

import (
	"crypto/rsa"
	"log"
	"net"
	"sync"
)

type Server struct {
	address           string
	clients           map[string]*Client
	connectionHandler func(net.Conn)
	listener          net.Listener
	running           bool
	mutex             sync.Mutex
	clientsMutex      sync.Mutex
}

type Client struct {
	id     string
	conn   net.Conn
	token  string
	pubKey *rsa.PublicKey
}

func NewServer(address string) (s *Server) {
	s = new(Server)
	s.address = address
	s.clients = map[string]*Client{}
	return
}

func (s *Server) Start() {
	var err error
	s.running = true
	s.listener, err = s.createListener()
	if err != nil {
		s.running = false
		log.Printf("Failed to listen on %v: %v", s.address, err)
		return
	}
	go s.listen()
}

func (s *Server) Stop() {
	if s.listener == nil {
		return
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.running = false
	if err := s.listener.Close(); err != nil {
		log.Printf("Could not close listener: %v", err)
	}
	for _, id := range s.GetClientIds() {
		s.CloseConnection(id)
	}
	s.listener = nil
}

func (s *Server) GetClient(id string) (*Client, bool) {
	s.clientsMutex.Lock()
	defer s.clientsMutex.Unlock()
	client, ok := s.clients[id]
	return client, ok
}

func (s *Server) PutClient(id string, client *Client) {
	s.clientsMutex.Lock()
	defer s.clientsMutex.Unlock()
	s.clients[id] = client
}

func (s *Server) GetClientIds() []string {
	s.clientsMutex.Lock()
	defer s.clientsMutex.Unlock()
	var clients []string
	for k := range s.clients {
		clients = append(clients, k)
	}
	return clients
}

func (s *Server) listen() {
	log.Print("Listening on ", s.address)

	for s.isRunning() {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Print("Could not accept connection: ", err)
		} else {
			go s.connectionHandler(conn)
		}
	}
}

func (s *Server) isRunning() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.running
}

func (s *Server) createListener() (net.Listener, error) {
	return net.Listen("tcp", s.address)
}

func (s *Server) CloseConnection(id string) {
	s.clientsMutex.Lock()
	defer s.clientsMutex.Unlock()
	delete(s.clients, id)
	log.Printf("Connection to %v closed", id)
}

func (c *Client) Ok() (reply *ControllerReply) {
	reply = new(ControllerReply)
	reply.StatusCode = new(ControllerReply_StatusCode)
	*reply.StatusCode = ControllerReply_OK
	return
}

func (c *Client) Reject(reason string) (reply *ControllerReply) {
	log.Print("Reject request: " + reason)
	reply = new(ControllerReply)
	reply.StatusCode = new(ControllerReply_StatusCode)
	*reply.StatusCode = ControllerReply_REJECTED
	reply.Reason = new(string)
	*reply.Reason = reason
	return
}
