package gApi

import (
	"database/sql"

	"github.com/vinayaknolastname/our/protobuf/video"
)

type GrpcServer struct {
	DB *sql.DB
	video.UnimplementedVideoServiceServer
}

func NewGApiServer(db *sql.DB) *GrpcServer {

	server := &GrpcServer{
		DB: db,
	}

	return server
}
