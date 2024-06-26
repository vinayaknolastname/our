package main

import (
	"log"
	"net"

	"github.com/vinayaknolastname/our/protobuf/video"
	"github.com/vinayaknolastname/our/services/common/db"
	"github.com/vinayaknolastname/our/services/common/utils"
	gApi "github.com/vinayaknolastname/our/services/video/gAPi"
	"google.golang.org/grpc"
)

func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Println("cofig load error  %err ", err)
	}

	db, err := db.NewDB()

	if err != nil {
		utils.LogSomething("db error", err, 0)
	}

	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", config.GrpcPortVideo)

	if err != nil {
		log.Println("tcp listenEroo  %err ", err)
	}

	server := gApi.NewGApiServer(db.Db)

	video.RegisterVideoServiceServer(grpcServer, server)

	grpcServer.Serve(listener)
}
