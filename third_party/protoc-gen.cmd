protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 todo-service.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --plugin=protoc-gen-grpc-gateway=$GOPATH/bin/protoc-gen-grpc-gateway --grpc-gateway_out=logtostderr=true:pkg/api/v1 todo-service.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --plugin=protoc-gen-swagger=$GOPATH/bin/protoc-gen-swagger --swagger_out=logtostderr=true:api/swagger/v1 todo-service.proto

protoc --proto_path=api/proto/v1 --proto_path=third_party --plugin=protoc-gen-grpc-gateway=$GOPATH/bin/protoc-gen-grpc-gateway --grpc-gateway_out=logtostderr=true:pkg/api/v1 todo-service.proto