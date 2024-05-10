package grpcHandlers

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/vinayaknolastname/our/gateway/utils"
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
	conn *grpc.ClientConn
}

var connection UserGrpcService

func ConnectUserServiceGrpc(c *gin.Context) {
	utils.LogSomething("Calling ConnectUserServiceGrpc", connection.conn, 1)

	if connection.conn == nil {
		conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
		if err != nil {
			utils.LogSomething("Connecting Grpc User Dial Err", err, 0)
		}
		connection = UserGrpcService{conn: conn}

	}

	utils.LogSomething("Connecting Grpc User Dial Err", connection.conn, 0)
	c.Next()
}

func CreateUser(c *gin.Context) {
	utils.LogSomething("dd", "sdd", 0)
}
