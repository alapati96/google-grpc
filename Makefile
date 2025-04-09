generate_grpc_code:
	protoc --go_out=demo --go_opt=paths=source_relative --go-grpc_out=demo --go-grpc_opt=paths=source_relative demo.proto