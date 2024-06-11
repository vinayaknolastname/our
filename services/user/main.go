package main

import (
	"log"
	"net"

	user "github.com/vinayaknolastname/our/protobuf/user"
	"github.com/vinayaknolastname/our/services/common/db"
	"github.com/vinayaknolastname/our/services/common/utils"
	gApi "github.com/vinayaknolastname/our/services/user/gapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Println("cofig load error  %err ", err)
	}
//D
	storage, err := db.NewDB()
	if err != nil {
		log.Println("db load error  %err ", err)
	}

	gRPCServer := grpc.NewServer()

	server, err := gApi.NewServer(*storage)

	user.RegisterUserServiceServer(gRPCServer, server)

	reflection.Register(gRPCServer)

	listener, err := net.Listen("tcp", config.GrpcPort)

	if err != nil {
		log.Fatal("Cannot Create GrpcServer %s", err)
	}

	log.Println("start grPC server at %s", listener.Addr().String())

	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot server gRPC")
	}

	//ffff
}
