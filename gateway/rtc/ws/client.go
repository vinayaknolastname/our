package ws

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
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
