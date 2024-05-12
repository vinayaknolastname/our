package ws

import (
	"log"

	grpcHandlers "github.com/vinayaknolastname/our/gateway/grpc"
	"github.com/vinayaknolastname/our/gateway/utils"
)

type WsManager struct {
	Chats      map[string]*Chat
	Register   chan *Client
	Unregister chan *Client
	Message    chan *Message
}

// var StoreWsManager *WsManager

func NewWsManager() *WsManager {
	return &WsManager{
		Chats:      make(map[string]*Chat),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Message:    make(chan *Message),
	}
}

func (w *WsManager) RunWsManager() {
	log.Println("w.Register")

	for {

		select {

		case cl := <-w.Register:
			log.Println("w.Register", cl)
			if _, ok := w.Chats[cl.ChatId]; ok {

				if _, ok := w.Chats[cl.ChatId].Clients[cl.ID]; ok == false {
					w.Chats[cl.ChatId].Clients[cl.ID] = cl

				}

				log.Println("w.Register", w.Chats)

				// if _, ok := r.Clients[cl.ID]; !ok {
				// 	r.Clients[cl.ID] = cl
				// }
			}
		// case cl := <-w.Unregister:
		// 	if _, ok := w.Clients[cl.ID]; ok == true {
		// 		delete(w.Clients, cl.ID)
		// 	}

		//  .Clients[cl.ID]; ok {

		// 	if len(h.Rooms[cl.RoomID].Clients) != 0 {

		// 		h.Broadcast <- &Message{
		// 			Content:  "User LEft",
		// 			RoomID:   cl.RoomID,
		// 			Username: cl.Username,
		// 		}
		// 	}
		// 	delete(h.Rooms[cl.RoomID].Clients, cl.ID)
		// 	close(cl.Message)

		case m := <-w.Message:
			utils.LogSomething("msg", w.Chats["2"].Clients, 1)
			if _, ok := w.Chats[m.ChatId]; ok {
				checkOtherUserIsConnectedOrNot()
				for _, cl := range w.Chats[m.ChatId].Clients {

					cl.Message <- m
				}
			}

		}
	}
}

func checkOtherUserIsConnectedOrNot(userId int32, chatId int32, clientsOfChat []*Client, content string) {
	for i := 0; i < len(clientsOfChat); i++ {
		if string(userId) == clientsOfChat[i].ID {
			grpcHandlers.CreateMessage(userId, chatId, content, true)
		}
	}
}

type Chat struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"client"`
}

// type Hub struct {
// 	Rooms      map[string]*Chat
// 	Register   chan *Clutils
// 	Unregister chan *Client
// 	Broadcast  chan *Message
// }

// func NewHub() *Hub {
// 	return &Hub{
// 		Rooms:      make(map[string]*Chat),
// 		Register:   make(chan *Client),
// 		Unregister: make(chan *Client),
// 		Broadcast:  make(chan *Message, 5),
// 	}
// }

// func (h *Hub) Run() {
// 	for {
// 		select {
// 		case cl := <-h.Register:
// 			if _, ok := h.Rooms[cl.ChatId]; ok {
// 				r := h.Rooms[cl.ChatId]
// 				if _, ok := r.Clients[cl.ID]; !ok {
// 					r.Clients[cl.ID] = cl
// 				}
// 			}
// 		case cl := <-h.Unregister:
// 			if _, ok := h.Rooms[cl.ChatId]; ok {
// 				if _, ok := h.Rooms[cl.ChatId].Clients[cl.ID]; ok {

// 					if len(h.Rooms[cl.ChatId].Clients) != 0 {

// 						h.Broadcast <- &Message{
// 							Content:  "User LEft",
// 							ChatId:   cl.ChatId,
// 							Username: cl.Username,
// 						}
// 					}
// 					delete(h.Rooms[cl.ChatId].Clients, cl.ID)
// 					close(cl.Message)
// 				}
// 			}

// 		case m := <-h.Broadcast:
// 			if _, ok := h.Rooms[m.ChatId]; ok {
// 				for _, cl := range h.Rooms[m.ChatId].Clients {
// 					cl.Message <- m
// 				}
// 			}

// 		}
// 	}
// }
