package grpcHandlers

import (
	context "context"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinayaknolastname/our/gateway/types"
	"github.com/vinayaknolastname/our/gateway/utils"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"
	"google.golang.org/grpc"
	// pb "google.golang.org/grpc/examples/route_guide/routeguide"
)

var opts []grpc.DialOption

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:9000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

type UserGrpcService struct {
	conn              *grpc.ClientConn
	userServiceClient *user.UserServiceClient
}

var connection UserGrpcService

func ConnectUserServiceGrpcMiddleWare(c *gin.Context) {
	utils.LogSomething("Calling ConnectUserServiceGrpc", connection.conn, 1)

	ConnectUserServiceGrpc()

	c.Next()
}

func ConnectUserServiceGrpc() {
	if connection.conn == nil {
		utils.LogSomething("Connection is nil Connecting user service", connection.conn, 1)

		conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
		if err != nil {
			utils.LogSomething("Connecting Grpc User Dial Err", err, 0)
		}
		connection = UserGrpcService{conn: conn}

	}

	utils.LogSomething("Connecting Grpc User Dial Err", "connection.conn", 0)
}

type CommonResponse struct {
	StatusCode int32  `json:"statusCode"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

func CreateUser(c *gin.Context) {
	utils.LogSomething("Create User Start", "", 1)

	name := c.Param("name")
	// intUserId, err := strconv.Atoi(userId)
	phone_number := c.Param("phone")
	intPhone, err := strconv.Atoi(phone_number)

	client := user.NewUserServiceClient(connection.conn)
	resp, err := client.CreateUser(context.Background(), &user.CreateUserRequest{Name: name, PhoneNumber: int32(intPhone)})
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

	var tempChats []types.Chat

	for i := 0; i < len(resp.UserData.Chat); i++ {

		chatData := types.Chat{
			ID:      resp.UserData.Chat[i].Id,
			Name:    resp.UserData.Chat[i].Name,
			Members: resp.UserData.Chat[i].Members,
		}
		err := append(tempChats, chatData)
		if err != nil {

		}
	}

	data := &types.UserAndChatData{
		UserId:       resp.UserData.Id,
		Phone_number: resp.UserData.PhoneNumber,
		Name:         resp.UserData.Name,
		Chats:        tempChats,
	}

	utils.LogSomething("Local saved user and chat", data, 1)

	c.JSON(int(resp.ResData.StatusCode), resp)
}

func GetUserAndChatsFunction(userGrpcService UserGrpcService, intUserId int32) types.UserAndChatData {

	if userGrpcService.conn == nil {
		ConnectUserServiceGrpc()
	}
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

	var tempChats []types.Chat

	for i := 0; i < len(resp.UserData.Chat); i++ {

		chatData := types.Chat{
			ID:      resp.UserData.Chat[i].Id,
			Name:    resp.UserData.Chat[i].Name,
			Members: resp.UserData.Chat[i].Members,
		}
		tempChats = append(tempChats, chatData)
		// if err != nil {
		// 	utils.LogSomething("resp.UserData.Chat", err, 0)
		// }
		println("temoChats", tempChats)
	}

	data := types.UserAndChatData{
		UserId:       resp.UserData.Id,
		Phone_number: resp.UserData.PhoneNumber,
		Name:         resp.UserData.Name,
		Chats:        tempChats,
	}

	return data
}

type StartChatRequest struct {
	Name     string  `json:"name"`
	ChatType int     `json:"type"`
	Members  []int32 `json:"members"`
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

func CreateMessage(userId int32, chatId int32, content string, isDelivered []int32) {
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

// func GetMessages(userId int32, chatId int32, sequence int32) {
// 	utils.LogSomething("Get Messages", "", 1)

// 	// userId := c.Param("userId")
// 	// intUserId, err := strconv.Atoi(userId)

// 	client := user.NewUserServiceClient(connection.conn)
// 	resp, err := client.GetMessages(context.Background(), &user.GetMessageRequest{UserId: userId, ChatId: chatId, Seq: 0})
// 	if err != nil {
// 		log.Fatalf("Failed to call Get Messages: %v", err)
// 	}

// 	log.Printf("Get Messages Success: %s", resp)

// 	utils.LogSomething("CreateMessage Success", resp, 1)

// }
