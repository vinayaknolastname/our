package grpcHandlers

import (
	context "context"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinayaknolastname/our/gateway/utils"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"
	"google.golang.org/grpc"
	// pb "google.golang.org/grpc/examples/route_guide/routeguide"
)

var opts []grpc.DialOption

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:61619", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

type UserGrpcService struct {
	conn              *grpc.ClientConn
	userServiceClient *user.UserServiceClient
}

var connection UserGrpcService

func ConnectUserServiceGrpc(c *gin.Context) {
	utils.LogSomething("Calling ConnectUserServiceGrpc", connection.conn, 1)

	if connection.conn == nil {
		utils.LogSomething("Connection is nil Connecting user service", connection.conn, 1)

		conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
		if err != nil {
			utils.LogSomething("Connecting Grpc User Dial Err", err, 0)
		}
		connection = UserGrpcService{conn: conn}

	}

	utils.LogSomething("Connecting Grpc User Dial Err", "connection.conn", 0)
	c.Next()
}

type CommonResponse struct {
	statusCode int32  `json:"statusCode"`
	success    bool   `json:"success"`
	message    string `json:"message"`
}

func CreateUser(c *gin.Context) {
	utils.LogSomething("Create User Start", "", 1)

	client := user.NewUserServiceClient(connection.conn)
	resp, err := client.CreateUser(context.Background(), &user.CreateUserRequest{Name: "Alice", PhoneNumber: 991881000})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	log.Printf("Response from server: %s", resp)
	// respto := CommonResponse{
	// 	statusCode: resp.ResData.StatusCode,
	// 	success:    resp.ResData.Success,
	// 	message:    resp.ResData.Message,
	// }

	utils.LogSomething("Grpc res into User", resp, 1)

	c.JSON(int(resp.ResData.StatusCode), resp.ResData)
}

func GetUserAndChats(c *gin.Context) {
	utils.LogSomething("Create User Start", "", 1)

	userId := c.Param("userId")
	intUserId, err := strconv.Atoi(userId)

	client := user.NewUserServiceClient(connection.conn)
	resp, err := client.GetUserData(context.Background(), &user.GetUserRequest{UserId: int32(intUserId)})
	if err != nil {
		log.Fatalf("Failed to call GetUserAndChats: %v", err)
	}

	log.Printf("Response from server: %s", resp)
	// respto := CommonResponse{
	// 	statusCode: resp.ResData.StatusCode,
	// 	success:    resp.ResData.Success,
	// 	message:    resp.ResData.Message,
	// }

	utils.LogSomething("Grpc res into getUserChats", resp, 1)

	c.JSON(int(resp.ResData.StatusCode), resp)
}

type StartChatRequest struct {
	Name     string   `json:"name"`
	ChatType int      `json:"type"`
	Members  []string `json:"members"`
}

func StartChat(c *gin.Context) {

	var req StartChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	utils.LogSomething("Started Start Chat Request", req, 1)

	client := user.NewUserServiceClient(connection.conn)
	resp, err := client.StartChat(context.Background(), &user.StartChatRequest{Name: req.Name, Type: int32(req.ChatType), Members: req.Members})
	if err != nil {
		log.Fatalf("Failed to call Start Chat: %v", err)
	}

	log.Printf("Response from server: %s", resp)
	// respto := CommonResponse{
	// 	statusCode: resp.ResData.StatusCode,
	// 	success:    resp.ResData.Success,
	// 	message:    resp.ResData.Message,
	// }

	utils.LogSomething("Grpc res into User", resp, 1)

	c.JSON(int(resp.StatusCode), resp)
}

func CreateMessage(userId int32, chatId int32, content string, isDelivered []int) {
	utils.LogSomething("Create Message", "", 1)

	// userId := c.Param("userId")
	// intUserId, err := strconv.Atoi(userId)

	client := user.NewUserServiceClient(connection.conn)
	resp, err := client.SendMessage(context.Background(), &user.CreateMessageRequest{UserId: userId, ChatId: chatId, Content: content, IsDelivered: isDelivered})
	if err != nil {
		log.Fatalf("Failed to call CreateMessage: %v", err)
	}

	log.Printf("CreateMessage Success: %s", resp)
	// respto := CommonResponse{
	// 	statusCode: resp.ResData.StatusCode,
	// 	success:    resp.ResData.Success,
	// 	message:    resp.ResData.Message,
	// }

	utils.LogSomething("CreateMessage Success", resp, 1)

	// c.JSON(int(resp.StatusCode), resp)
}

func GetMessages(userId int32, chatId int32, sequence int32) {
	utils.LogSomething("Get Messages", "", 1)

	// userId := c.Param("userId")
	// intUserId, err := strconv.Atoi(userId)

	client := user.NewUserServiceClient(connection.conn)
	resp, err := client.GetMessages(context.Background(), &user.GetMessageRequest{UserId: userId, ChatId: chatId, Seq: 0})
	if err != nil {
		log.Fatalf("Failed to call Get Messages: %v", err)
	}

	log.Printf("Get Messages Success: %s", resp)
	// respto := CommonResponse{
	// 	statusCode: resp.ResData.StatusCode,
	// 	success:    resp.ResData.Success,
	// 	message:    resp.ResData.Message,
	// }

	utils.LogSomething("CreateMessage Success", resp, 1)

	// c.JSON(int(resp.StatusCode), resp)
}
