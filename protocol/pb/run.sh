rm -rf *.go
#go grpc
protoc --go_out=plugins=grpc:. bucket.proto
#gateway
protoc --grpc-gateway_out=logtostderr=true:. bucket.proto