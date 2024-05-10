package gApi

import (
	"context"

	"github.com/vinayaknolastname/our/services/user/db"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"
	"github.com/vinayaknolastname/our/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *gAPI) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error) {

	query := db.CreateUserQuery()
	result, err := server.Db.Db.Exec(query, req.GetName(), req.GetPhoneNumber())

	if err != nil {
		utils.LogSomething("err in creating pg", err.Error(), 0)
	}

	utils.LogSomething("user created", result, 0)
	return &user.UserResponse{ResData: &user.CommonResponse{

		StatusCode: 200,
		Success:    true,
		Message:    "User Created",
	}, Data: &user.User{}}, status.Errorf(codes.Unimplemented, "method CreatePG not implemented %result ")
}
