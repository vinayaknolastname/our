package ws

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       int32  `json:id`
	ChatId   int32  `json:roomId`
	Username string `json:"username"`
}

type Message struct {
	Id          int32                `json:"id"`
	Content     string               `json:"content"`
	ChatId      int32                `json:"roomID"`
	Username    string               `json:"username"`
	SenderId    int32                `json:"senderId"`
	DateTime    *timestamp.Timestamp `json:"dateTime"`
	DeliveredTo []int32              `json:"deliveredTo"`
	ReadedBy    []int32              `json:"readedBy"`
	IsDeleted   bool                 `json:"isDeleted"`
	Seq         int32                `json:"seq"`
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}
		fmt.Println("msg", message)
		c.Conn.WriteJSON(message)

	}
}
func (c *Client) readMessage(h *WsManager) {
	// defer func() {
	// 	h.Unregister <- c
	// 	c.Conn.Close()
	// }()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("ws err", err)
			}
		}
		fmt.Println("readmsg", m)

		msg := &Message{
			Content:  string(m),
			ChatId:   c.ChatId,
			Username: c.Username,
		}
		fmt.Println("readmsg", msg)

		h.Message <- msg

	}
}
