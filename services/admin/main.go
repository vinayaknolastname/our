package main

// import (
// 	"fmt"
// 	"log"
// 	"net"

// 	gApi "github.com/vinayaknolastname/our/services/admin/gapi"
// 	"github.com/vinayaknolastname/our/services/common/admin"
// 	"github.com/vinayaknolastname/our/services/common/db"
// 	"github.com/vinayaknolastname/our/utils"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"
// )

// func main() {

// 	config, err := utils.LoadConfig(".")

// 	if err != nil {
// 		log.Println("cofig load error  %err ", err)
// 	}

// 	// portStr := fmt.Sprintf("%d", config.DBConfig.DBPort)

// 	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		config.DBConfig.DBHost, config.DBConfig.DBPort, config.DBConfig.DBUSER, config.DBConfig.DBPASSWORD, config.DBConfig.DBNAME)

// 	log.Println("db ddd %err ", postgresqlDbInfo)

// 	storage := db.ConnectDBFnc(postgresqlDbInfo)

// 	storage.DB.Close()

// 	gRPCServer := grpc.NewServer()

// 	server, err := gApi.NewServer(storage)

// 	admin.RegisterAdminServiceServer(gRPCServer, server)

// 	reflection.Register(gRPCServer)

// 	listener, err := net.Listen("tcp", config.GrpcPort)

// 	if err != nil {
// 		log.Fatal("Cannot Create GrpcServer %s", err)
// 	}

// 	log.Println("start grPC server at %s", listener.Addr().String())

// 	err = gRPCServer.Serve(listener)
// 	if err != nil {
// 		log.Fatal("cannot server gRPC")
// 	}

// }
