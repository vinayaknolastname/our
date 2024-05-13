package types

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/gorilla/websocket"
)

type Chat struct {
	ID      int32             `json:"id"`
	Name    string            `json:"name"`
	Clients map[int32]*Client `json:"client"`
	Members []int32           `json:"members"`
}

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

type UserAndChatData struct {
	UserId       int32
	Phone_number int32
	Name         string
	Chats        []Chat
}
