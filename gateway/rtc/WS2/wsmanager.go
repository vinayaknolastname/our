package ws

import "github.com/vinayaknolastname/our/gateway/rtc/ws"

type WsManager struct {
	Chats      map[string]*ws.Chat
	Register   chan *ws.Client
	Unregister chan *ws.Client
	Message    chan *ws.Message
}

// var StoreWsManager *WsManager

func NewWsManager() *WsManager {
	return &WsManager{
		Chats:      make(map[string]*ws.Chat),
		Register:   make(chan *ws.Client),
		Unregister: make(chan *ws.Client),
		Message:    make(chan *ws.Message),
	}
}

func (w *WsManager) RunWsManager() {

	for {
		select {
		case cl := <-w.Register:
			if _, ok := w.Chats[cl.ChatId]; ok == false {

				if _, ok := w.Chats[cl.ChatId].Clients[cl.ID]; ok == false {
					w.Chats[cl.ChatId].Clients[cl.ID] = cl

				}
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
			if _, ok := w.Chats[m.ChatId]; ok {

				for _, cl := range w.Chats[m.ChatId].Clients {

					cl.Message <- m
				}
			}

		}
	}
}
