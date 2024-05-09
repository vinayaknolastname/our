package grpcHandlers

import (
	"flag"

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

func ConnectUserServiceGrpc() {

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {

	}
	defer conn.Close()
}
