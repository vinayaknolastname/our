package main

import (
	"log"
	"net"

	"github.com/vinayaknolastname/our/services/user/db"
	gApi "github.com/vinayaknolastname/our/services/user/gapi"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"

	"github.com/vinayaknolastname/our/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Println("cofig load error  %err ", err)
	}

	// portStr := fmt.Sprintf("%d", config.DBConfig.DBPort)

	// postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	config.DBConfig.DBHost, config.DBConfig.DBPort, config.DBConfig.DBUSER, config.DBConfig.DBPASSWORD, config.DBConfig.DBNAME)

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

}
