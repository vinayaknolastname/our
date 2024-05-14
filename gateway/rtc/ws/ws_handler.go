package ws

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	grpcHandlers "github.com/vinayaknolastname/our/gateway/grpc"
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

	data := grpcHandlers.GetUserAndChatsFunction(grpcHandlers.UserGrpcService{}, int32(intClientID))

	messages := grpcHandlers.GetMessages(grpcHandlers.UserGrpcService{}, int32(intClientID), 0, int32(intCHatID))

	log.Println("messages", messages)

	for i := 0; i < len(data.Chats); i++ {
		if data.Chats[i].ID == int32(intCHatID) {
			h.manager.Chats[int32(intCHatID)].Members = data.Chats[i].Members
		}
	}

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
		SenderId: int32(intClientID),
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

}
