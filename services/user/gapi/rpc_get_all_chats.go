package gApi

import (
	"context"
	"net/http"

	"github.com/vinayaknolastname/our/services/user/db"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"

	"github.com/vinayaknolastname/our/utils"
)

func (server *gAPI) GetAllChats(ctx context.Context, req *user.GetReq) (*user.GetAllChatsResponse, error) {

	// var userData GetUserModel
	query := db.GetAllChatsQuery()
	result, err := server.Db.Db.Query(query)

	if err != nil {
		utils.LogSomething("err in getting chats", err, 0)
	}
	var chats []*user.Chats
	for result.Next() {
		var tempReac *user.Chats
		result.Scan(&tempReac.Id, &tempReac.Name, &tempReac.Type, &tempReac.Members, &tempReac.LastSeq)

		chats = append(chats, tempReac)
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
	query := db.GetAllChatsQuery()
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
