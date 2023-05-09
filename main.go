package nislib

//go:generate protoc --go_out=. data.proto
//go:generate protoc --go-grpc_out=require_unimplemented_servers=false:. data.proto
