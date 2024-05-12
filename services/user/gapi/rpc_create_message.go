package gApi

import (
	"context"
	"net/http"

	"github.com/vinayaknolastname/our/services/user/db"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"
	"github.com/vinayaknolastname/our/utils"
)

func (server *gAPI) SendMessage(ctx context.Context, req *user.CreateMessageRequest) (*user.CommonResponse, error) {

	query := db.CreateMessageQuery()
	result, err := server.Db.Db.Exec(query)

	if err != nil {
		utils.LogSomething("err in creating pg", err.Error(), 0)
	}

	utils.LogSomething("user created", result, 0)
	response := &user.CommonResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "User Created",
	}

	utils.LogSomething("user response", response, 1)

	// status.Errorf(codes.Unimplemented, "method CreateUser not implemented %r", err)
	return response, nil
}
