package grpcHandlers

import (
	context "context"
	"log"

	"github.com/vinayaknolastname/our/gateway/types"
	"github.com/vinayaknolastname/our/gateway/utils"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"
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

	utils.LogSomething("Create Reaction Seuccess", resp, 1)

}
