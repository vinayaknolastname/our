package gApi

import (
	"context"
	"net/http"
	"strconv"

	"github.com/lib/pq"
	"github.com/vinayaknolastname/our/services/user/db"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"
	"github.com/vinayaknolastname/our/utils"
)

func (server *gAPI) StartChat(ctx context.Context, req *user.StartChatRequest) (*user.CommonResponse, error) {

	query := db.CreateChatQuery()
	var id int

	utils.LogSomething("err in creating chat", req.GetName(), 1)

	result := server.Db.Db.QueryRow(query, req.GetName(), req.GetType(), pq.Array(req.GetMembers()))
	result.Scan(&id)
	err := result
	if err != nil {
		utils.LogSomething("err in creating chat", err, 1)
	}

	for i := 0; i < len(req.GetMembers()); i++ {
		var userId = req.GetMembers()[i]
		utils.LogSomething("Mebers", userId, 1)
		utils.LogSomething("ChatID", id, 1)

		userIdInt, err := strconv.Atoi(userId)

		if err != nil {
			utils.LogSomething("Mebers", userId, 1)

		}
		server.AddChatInUsersModel(userIdInt, id)
	}

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
