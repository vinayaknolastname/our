package gApi

import (
	"context"
	"net/http"
	"time"

	"github.com/vinayaknolastname/our/services/user/db"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"
	"github.com/vinayaknolastname/our/utils"
)

func (server *gAPI) SendMessage(ctx context.Context, req *user.CreateMessageRequest) (*user.CommonResponse, error) {

	query := db.CreateMessageQuery()
	result := server.Db.Db.QueryRow(query, req.GetContent(), req.GetChatId(), req.GetUserId(), time.Now(), req.GetIsDelivered(), req.GetIsDelivered(), false, 0)

	if result.Err() != nil {
		utils.LogSomething("err in creating message", result.Err(), 0)
	}

	utils.LogSomething("msg created", result, 0)
	response := &user.CommonResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "msg Created",
	}

	utils.LogSomething("user response", response, 1)

	// status.Errorf(codes.Unimplemented, "method CreateUser not implemented %r", err)
	return response, nil
}
