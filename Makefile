gen-proto-admin:
	@protoc	\
		--proto_path=protobuf "protobuf/admin.proto" \
		--go_out=services/common/admin --go_opt=paths=source_relative \
		--go-grpc_out=services/common/admin --go-grpc_opt=paths=source_relative \


////wlas uae linux
gen-proto-user:
	@protoc	\
		--proto_path=protobuf "protobuf/user.proto" \
		--go_out=services/user/proto_gen --go_opt=paths=source_relative \
		--go-grpc_out=services/user/proto_gen --go-grpc_opt=paths=source_relative \


start-server:
	go run services/admin/main.go


evans:
	evans --host localhost --port 9000 -r repl


docker-db:
	docker run --name ourdb -e POSTGRES_PASSWORD=ourdb -p 5432:5432 -v ourDB -d postgres

