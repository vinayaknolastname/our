package grpcHandlers

import (
	context "context"
	"log"

	// "github.com/gin-gonic/gin"
	"github.com/vinayaknolastname/our/gateway/types"
	"github.com/vinayaknolastname/our/gateway/utils"
	"github.com/vinayaknolastname/our/protobuf/user"
)

func CreateReaction(service UserGrpcService, reactionData types.ReactionOnMessageStruct) {
	utils.LogSomething("Create Reaction Start", "", 1)

	if service.conn == nil {
		ConnectUserServiceGrpc()
	}

	client := user.NewUserServiceClient(connection.conn)
	resp, err := client.ReactMessage(context.Background(), &user.SaveMessageReactionReq{MsgId: reactionData.MessageId, ReactorId: reactionData.ReactorId, ChatId: reactionData.ChatId, Reaction: reactionData.Reaction})
	if err != nil {
		log.Fatalf("Create Reaction Fail: %v", err)
	}

	// log.Printf("Create Reaction Success: %s", resp)

	utils.LogSomething("Create Reaction Success", resp, 1)

}

func GetMessages(userGrpcService UserGrpcService, userId int32, seq int32, chatId int32) interface{} {
	if userGrpcService.conn == nil {
		ConnectUserServiceGrpc()
	}

	client := user.NewUserServiceClient(connection.conn)
	resp, err := client.GetMessages(context.Background(), &user.GetMessageRequest{UserId: userId, ChatId: chatId, Seq: seq})
	if err != nil {
		log.Fatalf("Failed to call Get Messages: %v", err)
	}

	// log.Fatalf("Failed to call Get Messages: %v", resp)
	if len(resp.Message) > 0 {
		return resp.Message
	}

	return nil

}
