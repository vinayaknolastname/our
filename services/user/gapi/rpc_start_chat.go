package gApi

import (
	"context"
	"net/http"

	"github.com/lib/pq"
	user "github.com/vinayaknolastname/our/protobuf/user"
	"github.com/vinayaknolastname/our/services/common/utils"
	dbQ "github.com/vinayaknolastname/our/services/user/db"
)

func (server *gAPI) StartChat(ctx context.Context, req *user.StartChatRequest) (*user.CommonResponse, error) {

	query := dbQ.CreateChatQuery()
	var id int

	utils.LogSomething("err in creating chat", req.GetName(), 1)

	result := server.Db.Db.QueryRow(query, req.GetName(), req.GetType(), pq.Array(req.GetMembers()))
	result.Scan(&id)
	err := result
	if err != nil {
		utils.LogSomething("err in creating chat", err, 1)
	}

	listOfStringedMembers := req.GetMembers()
	// var wg sync.WaitGroup

	// responseFail := &user.CommonResponse{StatusCode: http.StatusOK,
	// 	Success: false,
	// 	Message: "chat not created",
	// }

	for i := 0; i < len(listOfStringedMembers); i++ {
		// wg.Add(1)
		var userId = listOfStringedMembers[i]
		utils.LogSomething("Mebers", userId, 1)
		utils.LogSomething("ChatID", id, 1)

		// userIdInt, err := strconv.Atoi(userId)

		// if err != nil {
		// 	utils.LogSomething("Mebers", userId, 1)
		// 	return responseFail, nil

		// }
		server.AddChatInUsersModel(int32(userId), int32(id))
	}

	// wg.Wait()

	// if result.Err() != nil {
	// 	utils.LogSomething("err in creating pg", result.Err(), 0)
	// }

	utils.LogSomething("chat created", result, 0)
	response := &user.CommonResponse{StatusCode: http.StatusOK,
		Success: true,
		Message: "chat created",
	}

	utils.LogSomething("user response", response, 1)

	// status.Errorf(codes.Unimplemented, "method CreateUser not implemented %r", err)
	return response, nil
}
