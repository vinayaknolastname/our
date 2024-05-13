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
	var id int32

	// var membersString string
	// members := req.GetIsDelivered()
	// for _, member := range members {
	// 	membersString += strconv.Itoa(int(member)) + ","
	// }
	// // Remove the trailing comma
	// membersString = strings.TrimSuffix(membersString, ",")
	result := server.Db.Db.QueryRow(query, req.GetContent(), req.GetChatId(), req.GetUserId(), time.Now(), false, 0)

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

	// status.Errorf(codes.Unimplemented, "method CreateUser not implemented %r", err)
	return response, nil
}

func AddDeliveryDataInDb(server *gAPI, messageId int32, deliveredTo int32) {
	query := db.AddDeliveredToInMessageProper()
	result, err := server.Db.Db.Exec(query, messageId, deliveredTo)

	if err != nil {
		utils.LogSomething("err in add delivered to data in db ", err, 0)

	}
	utils.LogSomething("err in add delivered to data in db", result, 0)

}
