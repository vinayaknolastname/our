package ws

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/vinayaknolastname/our/gateway/rtc/ws2"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:id`
	ChatId   string `json:roomId`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	ChatId   string `json:"roomID"`
	Username string `json:"username"`
}

func (c *Client) WriteMessage() {
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
func (c *Client) ReadMessage(h *ws2.WsManager) {
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

		msg := &Message{
			Content:  string(m),
			ChatId:   c.ChatId,
			Username: c.Username,
		}

		h.Message <- msg

	}
}
