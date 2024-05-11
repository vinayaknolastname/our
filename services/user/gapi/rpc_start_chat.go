package gApi

import (
	"context"
	"net/http"

	"github.com/vinayaknolastname/our/services/user/db"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"
	"github.com/vinayaknolastname/our/utils"
)

func (server *gAPI) StartChat(ctx context.Context, req *user.StartChatRequest) (*user.CommonResponse, error) {

	query := db.CreateChatQuery()
	var id int
	result := server.Db.Db.QueryRow(query, req.GetName(), req.GetMembers()).Scan(&id)

	err := result
	if err != nil {
		utils.LogSomething("err in creating chat", err, 1)
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
