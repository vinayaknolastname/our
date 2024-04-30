package gApi

import (
	"context"
	"fmt"

	adminDb "github.com/vinayakbot/our/services/admin/db"
	"github.com/vinayakbot/our/services/common/admin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *gAPI) CreatePG(ctx context.Context, req *admin.CreatePGRequest) (*admin.CommonResponse, error) {

	query := adminDb.AdminDB.CreatePgParams(adminDb.AdminDB{})
	result, err := server.storage.DB.Exec(query, req.GetName(), req.GetUserName(), req.GetPhoneNumber(), req.GetGender())

	if err != nil {
		fmt.Println("err in creating pg %e ", err)
	}

	fmt.Println("pg created %e ", result)
	return &admin.CommonResponse{StatusCode: 200, Success: true, Message: "ddd"}, status.Errorf(codes.Unimplemented, "method CreatePG not implemented")
}
