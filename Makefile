gen-proto-admin:
	@protoc	\
		--proto_path=protobuf "protobuf/admin.proto" \
		--go_out=services/common/admin --go_opt=paths=source_relative \
		--go-grpc_out=services/common/admin --go-grpc_opt=paths=source_relative \

gen-proto-user:
	@protoc	\
		--proto_path=protobuf "protobuf/user.proto" \
		--go_out=protobuf/user --go_opt=paths=source_relative \
		--go-grpc_out=protobuf/user --go-grpc_opt=paths=source_relative \

gen-proto-video:
	@protoc	\
		--proto_path=protobuf "protobuf/video.proto" \
		--go_out=protobuf/video --go_opt=paths=source_relative \
		--go-grpc_out=protobuf/video --go-grpc_opt=paths=source_relative \

start-server:
	go run services/admin/main.go


evans:
	evans --host localhost --port 9000 -r repl


docker-db:
	docker run --name ourdb -e POSTGRES_PASSWORD=ourdb -p 5432:5432 -v ourDB -d postgres


USER_BINARY=userServiceApp

build_user:
	@echo Building user binary...
	cd services/user && go build -o ${USER_BINARY} .
	ls ${USER_BINARY}  
	@echo Moving file..
	mv ${USER_BINARY} deploy/build
	@echo Done!

make_user_docker:
	docker buildx build -t your-image-name -f dockerFile .