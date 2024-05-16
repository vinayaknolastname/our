package grpcHandlers

import (
	context "context"
	"log"

	"github.com/gin-gonic/gin"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"
)

func GetAllChats(c *gin.Context) {

	client := user.NewUserServiceClient(connection.conn)
	resp, err := client.GetAllChats(context.Background(), &user.GetReq{Id: 0})
	if err != nil {
		log.Fatalf("Failed to call Get Messages: %v", err)
	}

	// log.Fatalf("Failed to call Get Messages: %v", resp)
	// if len(resp.Message) > 0 {
	// 	return resp.Message
	// }

	c.JSON(int(resp.ResData.StatusCode), resp)

}
