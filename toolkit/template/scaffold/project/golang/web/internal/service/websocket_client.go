package service

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 53 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 73 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	id string

	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// The processing function after receiving the data
	messageHandler func(c *Client, messageType int, p []byte)

	// Buffered channel of outbound messages.
	send     chan []byte
	sendJSON chan any

	afterClose func()
}

func NewClient(id string, hub *Hub, conn *websocket.Conn, messageHandler func(c *Client, messageType int, p []byte), afterClose func()) *Client {
	return &Client{
		id:             id,
		hub:            hub,
		conn:           conn,
		send:           make(chan []byte, 512),
		sendJSON:       make(chan any, 512),
		messageHandler: messageHandler,
		afterClose:     afterClose,
	}
}

func (c *Client) Run() {
	go c.ReadPump()
	go c.WritePump()
}

// ReadPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()

		if c.afterClose != nil {
			c.afterClose()
		}
	}()

	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		messageType, message, err := c.conn.ReadMessage()

		if err != nil {
			log.Printf("ws: id[%s] closed with error: %v", c.id, err)

			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				// do something
			}

			break
		}

		if c.messageHandler != nil {
			message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
			c.messageHandler(c, messageType, message)
		}
	}
}

// WritePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))

			if !ok {
				log.Printf("ws: hub[%s] closed client[%s].send channel", c.hub.id, c.id)
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err = w.Close(); err != nil {
				return
			}

		case message, ok := <-c.sendJSON:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))

			if !ok {
				log.Printf("ws: hub[%s] closed client[%s].send channel", c.hub.id, c.id)
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.conn.WriteJSON(message); err != nil {
				return
			}

			n := len(c.send)
			for i := 0; i < n; i++ {
				if err := c.conn.WriteJSON(<-c.send); err != nil {
					return
				}
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) WriteJSON(message any) {
	select {
	case c.sendJSON <- message:
	default:
		unregister(c)
	}
}

type Hub struct {
	id string

	// Registered clients.
	clients    map[*Client]bool
	idToClient map[string]*Client

	// Inbound messages from the clients.
	broadcast     chan []byte
	broadcastJSON chan any

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub(id string) *Hub {
	return &Hub{
		id:         id,
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		idToClient: make(map[string]*Client),
	}
}

func (h *Hub) ID() string {
	return h.id
}

func (h *Hub) Size() int {
	return len(h.clients)
}

func (h *Hub) BroadcastJSON(message any) {
	for client := range h.clients {
		select {
		case client.sendJSON <- message:
		default:
			unregister(client)
		}
	}
}

func (h *Hub) WriteJSON(id string, message any) error {
	client, err := h.findClient(id)
	if err != nil {
		return err
	}

	client.WriteJSON(message)
	return nil
}

func (h *Hub) findClient(id string) (*Client, error) {
	client, ok := h.idToClient[id]
	if !ok {
		return nil, fmt.Errorf("client[%s] not found", id)
	}
	return client, nil
}

func (h *Hub) ExistsClient(id string) bool {
	_, err := h.findClient(id)
	return err == nil
}

func (h *Hub) Register(c *Client) {
	h.register <- c
	log.Printf("ws: hub[%s] registered client[%s]", h.id, c.id)
}

func (h *Hub) Run() {
	defer func() {
		log.Printf("ws: hub[%s] exited", h.id)
	}()

	log.Printf("ws: hub[%s] started", h.id)

	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.idToClient[client.id] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				unregister(client)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					unregister(client)
				}
			}
		}
	}
}

type Scheduler struct {
	hubs       map[*Hub]bool
	masterHubs []*Hub
	idToHub    map[string]*Hub
	register   chan *Hub
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		hubs:     make(map[*Hub]bool),
		idToHub:  map[string]*Hub{},
		register: make(chan *Hub),
	}
}

func (s *Scheduler) Register(h *Hub, master bool) {
	s.register <- h

	if master {
		s.masterHubs = append(s.masterHubs, h)
	}

	log.Printf("ws: registered hub[%s]", h.id)
}

func (s *Scheduler) Run() {
	for {
		select {
		case hub := <-s.register:
			s.hubs[hub] = true
			s.idToHub[hub.id] = hub
		}
	}
}

func (s *Scheduler) findHub(id string) (*Hub, error) {
	hub, ok := s.idToHub[id]
	if !ok {
		return nil, fmt.Errorf("hub[%s] not found", id)
	}
	return hub, nil
}

func (s *Scheduler) FindOrCreateHub(id string, master bool) *Hub {
	if hub, ok := s.idToHub[id]; ok {
		return hub
	}
	return s.CreateHub(id, master)
}

func (s *Scheduler) CreateHub(id string, master bool) *Hub {
	if master && id != Default {
		if hub, err := s.findHub(Default); err == nil {
			hub.id = id
			s.idToHub[hub.id] = hub
			delete(s.idToHub, Default)
			s.CreateHub(Default, true)
			return hub
		}
	}

	hub := NewHub(id)
	go hub.Run()
	s.Register(hub, master)
	return hub
}

func (s *Scheduler) findClient(hubId, clientId string) (*Client, error) {
	hub, err := s.findHub(hubId)
	if err != nil {
		return nil, err
	}

	return hub.findClient(clientId)
}

func (s *Scheduler) FirstClient(clientId string) *Client {
	for i := range s.masterHubs {
		if client, err := s.masterHubs[i].findClient(clientId); err == nil {
			return client
		}
	}

	return nil
}

func (s *Scheduler) ExistsClient(hubId, clientId string) bool {
	_, err := s.findClient(hubId, clientId)
	return err == nil
}

func (s *Scheduler) BroadcastJSON(id string, message any) error {
	hub, err := s.findHub(id)
	if err != nil {
		return err
	}

	hub.BroadcastJSON(message)
	return nil
}

func (s *Scheduler) WriteJSON(hubId string, clientId string, message any) error {
	hub, err := s.findHub(hubId)
	if err != nil {
		return err
	}

	return hub.WriteJSON(clientId, message)
}

func (s *Scheduler) FindAvailableSlaveHub() *Hub {
	if len(s.masterHubs) == 0 {
		s.CreateHub(Default, true)
	}

	if len(s.masterHubs) == 1 {
		return s.masterHubs[0]
	}

	sort.Slice(s.masterHubs, func(i, j int) bool { return s.masterHubs[i].Size() < s.masterHubs[i].Size() })

	return s.masterHubs[0]
}

func unregister(client *Client) {
	delete(client.hub.clients, client)
	delete(client.hub.idToClient, client.id)
	close(client.send)

	log.Printf("ws: unregister client[%s]", client.id)
}

var CrossSiteUpgrader = websocket.Upgrader{
	HandshakeTimeout:  time.Second * 8,
	CheckOrigin:       func(r *http.Request) bool { return true }, // cross-site
	EnableCompression: false,
}
