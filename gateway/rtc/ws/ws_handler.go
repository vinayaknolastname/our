package ws

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{hub: h}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	fmt.Println("ss")

	var req CreateRoomReq

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}
	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	fmt.Println("Helleo")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID := c.Param("roomId")
	clientID := c.Query("userID")
	username := c.Query("username")

	cl := &Client{
		Conn:     conn,
		RoomID:   roomID,
		Username: username,
		ID:       clientID,
		Message:  make(chan *Message, 10),
	}

	m := &Message{
		Content:  "A new user joined",
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- cl

	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {

	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
	fmt.Println("Helleo")
	// conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, h.hub.Rooms)
	// roomID := c.Param("roomId")
	// clientID := c.Query("userID")
	// username := c.Query("username")

	// cl := &Client{
	// 	Conn:     conn,
	// 	RoomID:   roomID,
	// 	Username: username,
	// 	ID:       clientID,
	// 	Message:  make(chan *Message, 10),
	// }

	// m := &Message{
	// 	Content:  "A new user joined",
	// 	RoomID:   roomID,
	// 	Username: username,
	// }

	// h.hub.Register <- cl

	// h.hub.Broadcast <- m

	// go cl.writeMessage()
	// cl.readMessage(h.hub)
}
