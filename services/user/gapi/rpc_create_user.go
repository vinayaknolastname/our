package gApi

import (
	"context"
	"fmt"

	"github.com/vinayaknolastname/our/services/user/db"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *gAPI) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error) {

	query := db.CreateUserQuery()
	result, err := server.Db.Db.Exec(query, req.GetName(), req.GetUserName(), req.GetPhoneNumber(), "dsddd", req.GetPassword(), req.GetAuthToken())

	if err != nil {
		fmt.Println("err in creating pg %e ", err)
	}

	fmt.Println("pg created %e ", result)
	return &user.UserResponse{StatusCode: 200, Success: true, Message: "ddd"}, status.Errorf(codes.Unimplemented, "method CreatePG not implemented %result ")
}
