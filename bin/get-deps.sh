

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
GO111MODULE=off go get -v github.com/gogo/protobuf/protoc-gen-gogo
GO111MODULE=off go get -v github.com/envoyproxy/protoc-gen-validate
GO111MODULE=off go get -v github.com/amsokol/mongo-go-driver-protobuf
GO111MODULE=off go get -v github.com/amsokol/protoc-gen-gotagger
# GO111MODULE=off go get -v google.golang.org/protobuf/cmd/protoc-gen-go
GO111MODULE=on go get -v github.com/micro/micro/v2/cmd/protoc-gen-micro
GO111MODULE=on go get -u github.com/amsokol/mongo-go-driver-protobuf
GO111MODULE=on go get -u github.com/amsokol/protoc-gen-gotag
