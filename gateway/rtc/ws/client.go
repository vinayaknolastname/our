package ws

import (
	"encoding/json"
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

type Reaction struct {
	Id        int32  `json:"id"`
	ChatId    int32  `json:"chatId"`
	MessageId int32  `json:"messageId"`
	Reaction  string `json:"Reaction"`
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
	MsgType     string               `json:"msg_type"`
	MediaLink   string               `json:"media_link"`
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		fmt.Println("msg write", ok)

		if !ok {

			return
		}
		fmt.Println("msg", message)
		c.Conn.WriteJSON(message)

	}
}
func (c *Client) readMessage(h *WsManager) {
	defer func() {
		h.Unregister <- c
		c.Conn.Close()

	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("ws err", err)
			}
			return

		}

		formatedWebSocket := formatJsonWebSocketMessage(m)
		fmt.Println("Type:", formatedWebSocket.Type)
		// fmt.Println("Content:", msg.Content)
		fmt.Println("readmsg", m)
		msg := &Message{
			MsgType:   formatedWebSocket.Type,
			Content:   string(formatedWebSocket.Content),
			ChatId:    c.ChatId,
			Username:  c.Username,
			MediaLink: formatedWebSocket.MediaLink,
		}
		fmt.Println("readmsg", msg)

		h.Message <- msg

	}
}

type WsMessage struct {
	Type      string `json:"type"`
	Content   string `json:"content"`
	MediaLink string `json:"mediaLink"`
}

func formatJsonWebSocketMessage(data []byte) WsMessage {
	var msg WsMessage
	err := json.Unmarshal([]byte(data), &msg)
	if err != nil {
		fmt.Println("Error:", err)
		return msg
	}

	return msg

}
