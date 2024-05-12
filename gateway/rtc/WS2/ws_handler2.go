package ws

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/vinayaknolastname/our/gateway/rtc/ws"
	// "github.com/vinayaknolastname/our/gateway/rtc/ws"
)

type Handler struct {
	wsManager *WsManager
}

func NewHandler(h *WsManager) *Handler {
	return &Handler{wsManager: h}
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

	// h.hub.Rooms[req.ID] = &Chat{
	// 	ID:      req.ID,
	// 	Name:    req.Name,
	// 	Clients: make(map[string]*Client),
	// }
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
	clientID := c.Query("userID")
	username := c.Query("username")

	CheckChatIsLive(h, chatId)
	cl := &Client{
		Conn:     conn,
		ChatId:   chatId,
		Username: username,
		ID:       clientID,
		Message:  make(chan *ws.Message, 10),
	}

	m := &ws.Message{
		Content:  "Online",
		ChatId:   chatId,
		Username: username,
	}

	h.wsManager.Register <- cl

	h.wsManager.Message <- m

	go cl.WriteMessage()
	cl.ReadMessage(h.wsManager)
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func CheckChatIsLive(h *Handler, chatId string) {
	if _, ok := h.wsManager.Chats[chatId]; ok == false {
		h.wsManager.Chats[chatId] = &ws.Chat{
			ID:      chatId,
			Name:    chatId,
			Clients: make(map[string]*ws.Client),
		}
		return
	} else {
		return
	}
}

func (h *Handler) GetRooms(c *gin.Context) {

	rooms := make([]RoomRes, 0)

	for _, r := range h.wsManager.Chats {
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
