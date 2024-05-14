package gApi

import (
	"context"
	"errors"
	"net/http"

	"github.com/vinayaknolastname/our/services/user/db"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"
	"github.com/vinayaknolastname/our/utils"
)

func (server *gAPI) ReactMessage(ctx context.Context, req *user.SaveMessageReactionReq) (*user.CommonResponse, error) {

	query := db.CreateReactionQuery()
	result, err := server.Db.Db.Exec(query, req.GetReaction(), req.GetMsgId(), req.GetReactorId(), req.GetChatId())

	errorResponse := &user.CommonResponse{
		StatusCode: http.StatusBadRequest,
		Success:    false,
		Message:    "Reaction Not Created",
	}

	if err != nil {
		utils.LogSomething("err in creating react message", err.Error(), 0)
		return errorResponse, errors.New("err in creating react message")
	}

	utils.LogSomething("Reaction Created", result, 0)
	response := &user.CommonResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Reaction Created",
	}

	// status.Errorf(codes.Unimplemented, "method CreateUser not implemented %r", err)
	return response, nil
}
