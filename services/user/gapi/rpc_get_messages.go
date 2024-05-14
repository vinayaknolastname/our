package gApi

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/vinayaknolastname/our/services/user/db"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"

	"github.com/vinayaknolastname/our/utils"
)

type Message struct {
	Id           int32             `json:"id"`
	content      string            `json:"content"`
	senderId     int32             `json:"sender_id"`
	dateTime     time.Time         `json:"date_time"`
	deliveredTo  []int32           `json:"delivered_to"`
	readedBy     []int32           `json:"readed_by"`
	chatId       int32             `json:"chat_id"`
	isDeleted    bool              `json:"is_deleted`
	seq          int32             `json:"seq"`
	mediaLink    string            `json:"mediaLink"`
	msgReactions []MessageReaction `json:"reactions"`
}

type MessageReaction struct {
	Id        int32  `json:"id"`
	reaction  string `json:"reaction"`
	msgId     int32  `json:"msg_id"`
	reactorId int32  `json:"reactor_id"`
	chatId    int32  `json:"chat_id"`
}

func (server *gAPI) GetMessages(ctx context.Context, req *user.GetMessageRequest) (*user.MessageResponse, error) {

	var message []*user.Message

	chat := GetUserChats(server, req.GetChatId())

	for i := 1; i <= int(chat.seq); i++ {

		msg, err := GetMessage(server, int32(i), req.GetChatId())
		if err != nil {
			break
		}
		message = append(message, msg)
	}
	// query := db.GetMessageQuery()
	// msgResult, err := server.Db.Db.Query(query, req.ChatId)

	// if err != nil {
	// 	utils.LogSomething("err in get messages chats", err, 0)
	// 	return nil, errors.New("err in get messages chats ")
	// }

	// for msgResult.Next() {

	// }

	// errSCan := result.Scan(&chatModel.id, &chatModel.name, &chatModel.chatType, &chatModel.members)

	// if errSCan != nil {
	// 	utils.LogSomething("err in get user chats db:", errSCan, 0)
	// }

	resposne := &user.MessageResponse{
		ResData: &user.CommonResponse{
			StatusCode: http.StatusOK,
			Success:    true,
			Message:    "Messages fecthed",
		},
		Message: message,
	}

	return resposne, nil
}

func GetMessage(server *gAPI, seq int32, chatId int32) (*user.Message, error) {
	var message user.Message
	query := db.GetMessageQuery()
	msgResult := server.Db.Db.QueryRow(query, chatId, seq)

	if msgResult.Err() != nil {
		utils.LogSomething("err in get messages chats", msgResult.Err(), 0)
		return &user.Message{}, errors.New("err in get messages chats ")
	}

	msgResult.Scan(&message.Id, &message.Content, &message.ChatId, &message.SenderId,
		&message.DateTime, &message.DeliveredTo, &message.ReadedBy, &message.IsDeleted,
		&message.Seq, &message.MediaLink)

	reactions := GetReactions(server, message.Id)

	message.MsgReactions = reactions
	utils.LogSomething("message", message, 0)

	return &message, nil

}

func GetReactions(server *gAPI, msgId int32) []*user.MessageReaction {
	queryReac := db.GetReactionQuery()
	reacResult, err := server.Db.Db.Query(queryReac, msgId)
	if err != nil {
		utils.LogSomething("err in getting reactions", err, 0)
	}

	var reactions []*user.MessageReaction
	for reacResult.Next() {
		var tempReac user.MessageReaction
		reacResult.Scan(&tempReac.Id, &tempReac.Reaction, &tempReac.MsgId, &tempReac.ReactorId, &tempReac.ChatId)

		reactions = append(reactions, &tempReac)
	}

	return reactions
}
