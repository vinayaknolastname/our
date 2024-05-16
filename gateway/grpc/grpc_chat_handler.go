package grpcHandlers

import (
	context "context"
	"log"

	user "github.com/vinayaknolastname/our/services/user/proto_gen"
)

func GetAllChats(userGrpcService UserGrpcService) interface{} {
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
