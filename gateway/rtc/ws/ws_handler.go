package ws

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	manager *WsManager
}

func NewHandler(h *WsManager) *Handler {
	return &Handler{manager: h}
}

type CreateRoomReq struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	fmt.Println("ss")

	var req CreateRoomReq

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	h.manager.Chats[req.ID] = &Chat{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[int32]*Client),
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

func (h *Handler) StartChat(c *gin.Context) {
	fmt.Println("Helleo")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chatId := c.Param("chatId")
	intCHatID, err := strconv.Atoi(chatId)
	clientID := c.Query("userID")
	intClientID, err := strconv.Atoi(clientID)

	username := c.Query("username")
	CheckChatIsLive(h, int32(intCHatID))
	log.Println("chats", h.manager.Chats)
	cl := &Client{
		Conn:     conn,
		ChatId:   int32(intCHatID),
		Username: username,
		ID:       int32(intClientID),
		Message:  make(chan *Message, 10),
	}
	log.Println("chats", cl)
	m := &Message{
		Content:  "A new user joined",
		ChatId:   int32(intCHatID),
		Username: username,
	}
	log.Println("chats", m)
	h.manager.Register <- cl

	h.manager.Message <- m

	go cl.writeMessage()
	cl.readMessage(h.manager)
}

func CheckChatIsLive(h *Handler, chatId int32) {
	if _, ok := h.manager.Chats[chatId]; ok == false {
		h.manager.Chats[chatId] = &Chat{
			ID:      chatId,
			Name:    string(chatId),
			Clients: make(map[int32]*Client),
		}
		return
	} else {
		return
	}
}

type RoomRes struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {

	rooms := make([]RoomRes, 0)

	for _, r := range h.manager.Chats {
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
