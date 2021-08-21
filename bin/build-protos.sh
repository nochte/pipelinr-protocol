rm -rf protobuf/*
protoc \
        -I /usr/local/include \
        -I ${HOME}/go/1.13.0/src/github.com/gogo/protobuf/gogoproto/ \
        -I protocol/accounts \
        -I protocol/messages \
        -I protocol/pipes \
        -I protocol/healthcheck \
        -I ${HOME}/go/1.13.0/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/go/1.13.0/src/github.com/envoyproxy/protoc-gen-validate \
        -I ${HOME}/go/1.13.0/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/go/1.13.0/src/github.com/amsokol/protoc-gen-gotagger/proto \
        protocol/messages/messageenvelop.proto \
        protocol/messages/event.proto \
        --micro_out=./protobuf \
        --gogo_out=M:./protobuf \
        --validate_out="lang=go:./protobuf"


protoc \
        -I /usr/local/include \
        -I ${HOME}/go/1.13.0/src/github.com/gogo/protobuf/gogoproto/ \
        -I protocol/accounts \
        -I protocol/messages \
        -I protocol/pipes \
        -I protocol/healthcheck \
        -I ${HOME}/go/1.13.0/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/go/1.13.0/src/github.com/envoyproxy/protoc-gen-validate \
        -I ${HOME}/go/1.13.0/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/go/1.13.0/src/github.com/amsokol/protoc-gen-gotagger/proto \
        protocol/pipes/pipe.proto \
        --micro_out=./protobuf \
        --gogo_out=M:./protobuf \
        --validate_out="lang=go:./protobuf"

protoc \
        -I /usr/local/include \
        -I ${HOME}/go/1.13.0/src/github.com/gogo/protobuf/gogoproto/ \
        -I protocol/accounts \
        -I protocol/messages \
        -I protocol/pipes \
        -I protocol/healthcheck \
        -I ${HOME}/go/1.13.0/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/go/1.13.0/src/github.com/envoyproxy/protoc-gen-validate \
        -I ${HOME}/go/1.13.0/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/go/1.13.0/src/github.com/amsokol/protoc-gen-gotagger/proto \
        protocol/healthcheck/healthcheck.proto \
        --micro_out=./protobuf \
        --gogo_out=M:./protobuf \
        --validate_out="lang=go:./protobuf"

protoc \
        -I /usr/local/include \
        -I ${HOME}/go/1.13.0/src/github.com/gogo/protobuf/gogoproto/ \
        -I protocol/messages \
        -I protocol/pipes \
        -I protocol/healthcheck \
        -I protocol/accounts \
        -I ${HOME}/go/1.13.0/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/go/1.13.0/src/github.com/envoyproxy/protoc-gen-validate \
        -I ${HOME}/go/1.13.0/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/go/1.13.0/src/github.com/amsokol/protoc-gen-gotagger/proto \
        protocol/accounts/user.proto \
        --micro_out=./protobuf \
        --gogo_out=M:./protobuf \
        --validate_out="lang=go:./protobuf"

mv protobuf/github.com/nochte/pipelinr/protobuf/* protobuf/
rm -rf protobuf/github.com


# below this for mac / gvm
        # -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        # -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/envoyproxy/protoc-gen-validate \
        # -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        # -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/protoc-gen-gotagger/proto \
