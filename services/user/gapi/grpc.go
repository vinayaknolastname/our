package gApi

import (
	"github.com/vinayaknolastname/our/services/user/db"
	user "github.com/vinayaknolastname/our/services/user/proto_gen"
)

type gAPI struct {
	user.UnimplementedUserServiceServer
	Db *db.Database
}

// func NewGRPCServer(addr string) *gRPCServer {
// 	return &gRPCServer{addr: addr}
// }

func NewServer(db db.Database) (*gAPI, error) {

	server := &gAPI{
		Db: &db,
	}

	return server, nil
	// lis, err := net.Listen("tcp", s.addr)

	// if err != nil {
	// 	log.Fatalf("Failed to listen", err)
	// }

}
