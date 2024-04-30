gen:
	@protoc	\
		--proto_path=protobuf "protobuf/admin.proto" \
		--go_out=services/common/admin --go_opt=paths=source_relative \
		--go-grpc_out=services/common/admin --go-grpc_opt=paths=source_relative \

start-server:
	go run services/admin/main.go


evans:
	evans --host localhost --port 9000 -r repl


docker-admin:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=ourDB -e POSTGRES_PASSWORD=ourDB -v ourDB