package grpcHandlers

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vinayaknolastname/our/gateway/utils"
	"github.com/vinayaknolastname/our/protobuf/video"
	"google.golang.org/grpc"
)

var (
	// tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddrVideo = flag.String("addr", "localhost:9001", "The server address in the format of host:port")
	// serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

type VideoGrpcService struct {
	conn               *grpc.ClientConn
	videoServiceClient video.VideoServiceClient
}

var videoGrpcClient *VideoGrpcService

func ConnectWithVideoGrpcMiddleWare(c *gin.Context) {
	utils.LogSomething("Calling ConnectVideoServiceGrpc --- Middle", connection.conn, 1)

	ConnectWithVideoGrpc()
	c.Next()
}

func ConnectWithVideoGrpc() {

	if videoGrpcClient.conn == nil {

		connection, err := grpc.NewClient(*serverAddrVideo, grpc.WithInsecure())

		if err != nil {
			utils.LogSomething("err in dialing video grpc", err, 0)
		}

		client := video.NewVideoServiceClient(connection)

		videoGrpcClient = &VideoGrpcService{
			conn:               connection,
			videoServiceClient: client,
		}
	}
}

func StartVideoStream(c *gin.Context) {

	videoData := &video.Video{
		Content: []byte{111},
	}

	client := videoGrpcClient.videoServiceClient

	stream, err := client.VideoProccess(context.Background())

	if err != nil {
		utils.LogSomething("videoProces", err, 0)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("failed to receive a message : %v", err)
			}
			log.Printf("Received message from %s: %s", res.StatusCode, res.Message)
		}
	}()

	// users := []string{"Alice", "Bob", "Charlie"}
	go func() {
		for {
			if err := stream.Send(videoData); err != nil {
				log.Fatalf("failed to send a message: %v", err)
			}
			time.Sleep(1 * time.Second)
		}

	}()
	stream.CloseSend()
	<-waitc

}
