package gapi

import "database/sql"

type GrpcServer struct {
	DB *sql.DB
}
