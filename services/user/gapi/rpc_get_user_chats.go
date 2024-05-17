package gApi

import (
	"context"
	"net/http"

	"github.com/lib/pq"
	user "github.com/vinayaknolastname/our/protobuf/user"
	"github.com/vinayaknolastname/our/services/common/utils"
	dbQ "github.com/vinayaknolastname/our/services/user/db"
)

type GetUserModel struct {
	id           int32
	name         string
	phone_number int32
	chats        pq.Int32Array
}

func (server *gAPI) GetUserData(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {

	var userData GetUserModel
	query := dbQ.GetUserQuery()
	result := server.Db.Db.QueryRow(query, req.UserId)

	if result.Err() != nil {
		utils.LogSomething("err in creating pg", result.Err(), 0)
	}

	errSCan := result.Scan(&userData.id, &userData.name, &userData.phone_number, &userData.chats)

	if errSCan != nil {
		utils.LogSomething("err in creating pg", errSCan, 0)
	}
	utils.LogSomething("err in creating pg", userData.id, 0)

	var userChats []*user.Chats
	for i := 0; i < len(userData.chats); i++ {
		chat := GetUserChats(server, userData.chats[i])

		grpcChat := user.Chats{
			Id:      chat.id,
			Name:    chat.name,
			Type:    chat.chatType,
			Members: chat.members,
		}
		userChats = append(userChats, &grpcChat)
	}

	utils.LogSomething("user created", userData, 0)
	response := &user.GetUserResponse{ResData: &user.CommonResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "User Created",
	},
		UserData: &user.User{
			Id:          userData.id,
			Name:        userData.name,
			PhoneNumber: userData.phone_number,
			Chat:        userChats,
		},
	}

	utils.LogSomething("user response", response, 1)

	// status.Errorf(codes.Unimplemented, "method CreateUser not implemented %r", err)
	return response, nil
}

type ChatModel struct {
	id       int32
	name     string
	chatType int32
	members  pq.Int32Array
	seq      int32
}

func GetUserChats(server *gAPI, chatId int32) ChatModel {

	var chatModel ChatModel
	query := dbQ.GetChatRowQuery()
	result := server.Db.Db.QueryRow(query, chatId)

	if result.Err() != nil {
		utils.LogSomething("err in get user chats db:", result.Err(), 0)
	}

	errSCan := result.Scan(&chatModel.id, &chatModel.name, &chatModel.chatType, &chatModel.members, &chatModel.seq)

	if errSCan != nil {
		utils.LogSomething("err in get user chats db:", errSCan, 0)
	}

	return chatModel
}
