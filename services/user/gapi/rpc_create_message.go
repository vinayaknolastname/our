package gApi

import (
	"context"
	"net/http"
	"time"

	user "github.com/vinayaknolastname/our/protobuf/user"
	"github.com/vinayaknolastname/our/services/common/utils"
	dbQ "github.com/vinayaknolastname/our/services/user/db"
)

func (server *gAPI) SendMessage(ctx context.Context, req *user.CreateMessageRequest) (*user.CommonResponse, error) {

	chat := GetUserChats(server, req.GetChatId())

	newSeq := chat.seq + 1

	query := dbQ.CreateMessageQuery()

	var id int32

	result := server.Db.Db.QueryRow(query, req.GetContent(), req.GetChatId(), req.GetUserId(), time.Now(), false, int32(newSeq), req.GetMediaLink())

	responseBad := &user.CommonResponse{
		StatusCode: http.StatusBadRequest,
		Success:    false,
		Message:    "Message not Created",
	}

	if result.Err() != nil {
		utils.LogSomething("Err in creating message DB", result.Err(), 0)

		return responseBad, nil
	}
	result.Scan(&id)

	listOfDeliveredTo := req.GetIsDelivered()
	for i := 0; i < len(listOfDeliveredTo); i++ {
		AddDeliveryDataInDb(server, id, int32(listOfDeliveredTo[i]))
	}

	utils.LogSomething("msg created", result, 0)
	response := &user.CommonResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "msg Created",
	}

	utils.LogSomething("user response", response, 1)

	queryUpdateSeq := dbQ.UpdateSeqInChat()
	_, err := server.Db.Db.Exec(queryUpdateSeq, newSeq, req.GetChatId())
	if err != nil {
		utils.LogSomething("errro in update seq", err, 0)
		return nil, err
	}
	// status.Errorf(codes.Unimplemented, "method CreateUser not implemented %r", err)
	return response, nil
}

func AddDeliveryDataInDb(server *gAPI, messageId int32, deliveredTo int32) {
	query := dbQ.AddDeliveredToInMessageProper()
	result, err := server.Db.Db.Exec(query, messageId, deliveredTo)

	if err != nil {
		utils.LogSomething("err in add delivered to data in db ", err, 0)

	}
	utils.LogSomething("err in add delivered to data in db", result, 0)

}
