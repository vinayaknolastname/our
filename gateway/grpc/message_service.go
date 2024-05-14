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
	resp, err := client.CreateUser(context.Background(), &user. {Name: name, PhoneNumber: int32(intPhone)})
	if err != nil {
		log.Fatalf("User Created: %v", err)
	}

	log.Printf("User Created: %s", resp)
	// respto := CommonResponse{
	// 	statusCode: resp.ResData.StatusCode,
	// 	success:    resp.ResData.Success,
	// 	message:    resp.ResData.Message,
	// }

	utils.LogSomething("User Created  User", resp, 1)

	c.JSON(int(resp.ResData.StatusCode), resp.ResData)
}
