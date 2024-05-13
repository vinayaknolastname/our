package ws

import (
	"log"

	// grpcHandlers "github.com/vinayaknolastname/our/gateway/grpc"
	// grpcHandlers "github.com/vinayaknolastname/our/gateway/grpc"
	grpcHandlers "github.com/vinayaknolastname/our/gateway/grpc"
	"github.com/vinayaknolastname/our/gateway/utils"
)

type WsManager struct {
	Chats      map[int32]*Chat
	Register   chan *Client
	Unregister chan *Client
	Message    chan *Message
}

// var StoreWsManager *WsManager

func NewWsManager() *WsManager {
	return &WsManager{
		Chats:      make(map[int32]*Chat),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Message:    make(chan *Message),
	}
}

func (w *WsManager) RunWsManager() {
	log.Println("w.Register")
	go broad(w)
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
		case cl := <-w.Unregister:
			if _, ok := w.Chats[cl.ChatId]; ok {
				if _, ok := w.Chats[cl.ChatId].Clients[cl.ID]; ok {

					// if len(w.Chats[cl.ChatId].Clients) != 0 {

					// }
					delete(w.Chats[cl.ChatId].Clients, cl.ID)
					close(cl.Message)
					m := &Message{
						Content:  "A guy left",
						ChatId:   int32(cl.ChatId),
						Username: "username",
						SenderId: int32(cl.ID),
					}
					utils.LogSomething("unbre", m, 1)

					w.Message <- m
				}
			}

			// case m := <-w.Message:
			// 	utils.LogSomething("msgHub", w.Chats[m.ChatId].Clients, 1)
			// 	if _, ok := w.Chats[m.ChatId]; ok {
			// 		membersOfChat := w.Chats[m.ChatId].Members
			// 		checkOtherUserIsConnectedOrNot(membersOfChat, m.ChatId, w.Chats[m.ChatId].Clients, m.Content, m.SenderId)
			// 		for _, cl := range w.Chats[m.ChatId].Clients {

			// 			cl.Message <- m
			// 		}
			// 	}

		}
	}

}

func broad(w *WsManager) {
	for {
		m := <-w.Message
		utils.LogSomething("msgHub", w.Chats[m.ChatId].Clients, 1)
		if _, ok := w.Chats[m.ChatId]; ok {
			membersOfChat := w.Chats[m.ChatId].Members
			checkOtherUserIsConnectedOrNot(membersOfChat, m.ChatId, w.Chats[m.ChatId].Clients, m.Content, m.SenderId)
			for _, cl := range w.Chats[m.ChatId].Clients {

				cl.Message <- m
			}
		}
	}
}

func checkOtherUserIsConnectedOrNot(membersOfChat []int32, chatId int32, clientsOfChat map[int32]*Client, content string, userId int32) {
	var tempDeliveredList []int32
	for i := 0; i < len(membersOfChat); i++ {
		if _, ok := clientsOfChat[membersOfChat[i]]; ok {

			tempDeliveredList = append(tempDeliveredList, membersOfChat[i])
		}

	}
	grpcHandlers.CreateMessage(userId, chatId, content, tempDeliveredList)

}

type Chat struct {
	ID      int32             `json:"id"`
	Name    string            `json:"name"`
	Clients map[int32]*Client `json:"client"`
	Members []int32           `json:"members"`
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
