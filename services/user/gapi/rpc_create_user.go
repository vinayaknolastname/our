package gApi

import (
	"context"
	"net/http"

	user "github.com/vinayaknolastname/our/protobuf/user"
	"github.com/vinayaknolastname/our/services/common/utils"
	dbQ "github.com/vinayaknolastname/our/services/user/db"
)

func (server *gAPI) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error) {

	query := dbQ.CreateUserQuery()
	result := server.Db.Db.QueryRow(query, req.GetName(), req.GetPhoneNumber())

	if result.Err() != nil {
		utils.LogSomething("err in creating pg", result.Err(), 0)
	}

	var userData user.User

	result.Scan(&userData.Id, &userData.Name, &userData.PhoneNumber, &userData.Chat)
	utils.LogSomething("user created", result, 0)
	response := &user.UserResponse{ResData: &user.CommonResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "User Created",
	},
		UserData: &userData,
	}

	utils.LogSomething("user response", response, 1)

	// status.Errorf(codes.Unimplemented, "method CreateUser not implemented %r", err)
	return response, nil
}

func (server *gAPI) AddChatInUsersModel(userId, chatId int32) {
	// mutex.Lock()
	// query := db.AddChatInUser()
	// var chats []uint8
	// err := server.Db.Db.QueryRow(query, userId).Scan(&chats)
	// if err != nil {
	// 	utils.LogSomething("err in add chat in Users Model query", err.Error(), 0)
	// }

	// if len(chats) == 0 {
	queryPro := dbQ.AddChatInUserProper()
	res, err := server.Db.Db.Exec(queryPro, userId, chatId)
	if err != nil {
		utils.LogSomething("err in add chat in Users Model when no data before err", err, 0)

	}
	utils.LogSomething("err in add chat in Users Model when no data before res", res, 0)
	// }

	// utils.LogSomething("err in add chat in Users Model", chats, 0)

	// chatsStr := string(chats)

	// chatsStr = strings.Trim(chatsStr, "{}")

	// // Split the string by comma to get individual values
	// chatIDsStr := strings.Split(chatsStr, ",")

	// // Convert each string value to an integer
	// var chatIDs []int
	// for _, chatIDStr := range chatIDsStr {
	// 	chatID, err := strconv.Atoi(strings.TrimSpace(chatIDStr))
	// 	if err != nil {
	// 		fmt.Println("Error converting chat ID:", err)
	// 		continue
	// 	}
	// 	chatIDs = append(chatIDs, chatID)
	// }
	// // utils.LogSomething("err in add chat in Users Model", chatsStr, 0)

	// chatIDs = append(chatIDs, chatId)

	// queryPro := db.AddChatInUserProper()
	// utils.LogSomething("err in add chat in Users Model", chatIDs, 0)

	// res, err2 := server.Db.Db.Exec(queryPro, chatIDs)
	// if err2 != nil {
	// 	utils.LogSomething("err in add chat in Users Model query", err.Error(), 0)

	// }
	// // // var chatIDs []int
	// // if err := json.Unmarshal([]byte(chatsStr), &chatIDs); err != nil {
	// // 	utils.LogSomething("Error:", err, 0)
	// // } else {
	// utils.LogSomething("Chat IDs:", queryPro, 1)
	// }

	// if result != nil {
	// 	utils.LogSomething("err in add chat in Users Model", result, 0)
	// }
	// mutex.Unlock()

}
