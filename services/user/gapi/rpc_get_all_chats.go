package gApi

import (
	"context"
	"net/http"

	"github.com/lib/pq"
	user "github.com/vinayaknolastname/our/protobuf/user"
	"github.com/vinayaknolastname/our/services/common/utils"
	dbQ "github.com/vinayaknolastname/our/services/user/db"
)

func (server *gAPI) GetAllChats(ctx context.Context, req *user.GetReq) (*user.GetAllChatsResponse, error) {

	// var userData GetUserModel
	query := dbQ.GetAllChatsQuery()
	result, err := server.Db.Db.Query(query)

	if err != nil {
		utils.LogSomething("err in getting chats", err, 0)
	}

	utils.LogSomething("err in getting chats", result, 0)

	var chats []*user.Chats
	for result.Next() {
		var tempReac user.Chats
		var member pq.Int32Array
		result.Scan(&tempReac.Id, &tempReac.Name, &tempReac.Type, &member, &tempReac.LastSeq)

		tempReac.Members = member
		chats = append(chats, &tempReac)
	}
	// errSCan := result.Scan(&userData.id, &userData.name, &userData.phone_number, &userData.chats)

	// if errSCan != nil {
	// 	utils.LogSomething("err in creating pg", errSCan, 0)
	// }
	// utils.LogSomething("err in creating pg", userData.id, 0)

	// var userChats []*user.Chats
	// for i := 0; i < len(userData.chats); i++ {
	// 	chat := GetUserChats(server, userData.chats[i])

	// 	grpcChat := user.Chats{
	// 		Id:      chat.id,
	// 		Name:    chat.name,
	// 		Type:    chat.chatType,
	// 		Members: chat.members,
	// 	}
	// 	userChats = append(userChats, &grpcChat)
	// }

	// utils.LogSomething("user created", userData, 0)
	response := &user.GetAllChatsResponse{ResData: &user.CommonResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "User Created",
	},
		Chat: chats}

	utils.LogSomething("user response", response, 1)

	// status.Errorf(codes.Unimplemented, "method CreateUser not implemented %r", err)
	return response, nil
}

func GetAllChats(server *gAPI, chatId int32) ChatModel {

	var chatModel ChatModel
	query := dbQ.GetAllChatsQuery()
	result, err := server.Db.Db.Query(query, chatId)

	if err != nil {
		utils.LogSomething("err in get user chats db:", result.Err(), 0)
	}

	errSCan := result.Scan(&chatModel.id, &chatModel.name, &chatModel.chatType, &chatModel.members, &chatModel.seq)

	if errSCan != nil {
		utils.LogSomething("err in get user chats db:", errSCan, 0)
	}

	return chatModel
}
