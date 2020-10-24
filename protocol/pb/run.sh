rm -rf *.go
#go grpc
protoc --go_out=plugins=grpc:. service.proto
#gateway
protoc --grpc-gateway_out=logtostderr=true:. service.proto