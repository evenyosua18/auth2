SEEDER_NAME?=

setup:
	@cp .env.example app/infrastructure/server/grpc/.env

clean:
	@echo "--- cleanup all build and generated files ---"
	@rm -vf app/infrastructure/proto/pb/*.pb.go

protoc: clean
	@echo "--- preparing proto output directories ---"
	@mkdir -p app/infrastructure/proto/pb

	@echo "--- Compiling all proto files ---"
#	@cd ./app/infrastructure/proto && protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative *.proto
	@cd ./app/infrastructure/proto && protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative *.proto

build-rest: protoc
	@echo "--- Building binary file ---"
	@go build -o ./main cmd/rest/rest.go

build-grpc: protoc
	@echo "--- Building binary file ---"
	@go build -o ./main cmd/grpc/grpc.go

run-grpc:
	@go run cmd/grpc/grpc.go

seed:
	@go run cmd/seeder/main.go $(SEEDER_NAME)

wire:
	@echo "-- generate dependency injection --"
	@wire ./app/infrastructure/container/