package gApi

import (
	"context"
	"net/http"

	"github.com/vinayaknolastname/our/services/user/db"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"
	"github.com/vinayaknolastname/our/utils"
)

func (server *gAPI) StartChat(ctx context.Context, req *user.StartChatRequest) (*user.CommonResponse, error) {

	query := db.CreateUserQuery()
	result, err := server.Db.Db.Exec(query, req.GetName(), req.GetPhoneNumber())

	if err != nil {
		utils.LogSomething("err in creating pg", err.Error(), 0)
	}

	utils.LogSomething("user created", result, 0)
	response := &user.UserResponse{ResData: &user.CommonResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "User Created",
	}}

	utils.LogSomething("user response", response, 1)

	// status.Errorf(codes.Unimplemented, "method CreateUser not implemented %r", err)
	return response, nil
}
