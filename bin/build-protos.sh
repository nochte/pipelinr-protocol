rm -rf protobuf/*
rm js/protobuf/*.js

protoc \
        -I /usr/local/include \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/gogo/protobuf/gogoproto/ \
        -I protocol/accounts \
        -I protocol/messages \
        -I protocol/pipes \
        -I protocol/healthcheck \
        -I protocol/accounting \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/envoyproxy/protoc-gen-validate \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/protoc-gen-gotagger/proto \
        protocol/messages/messageenvelop.proto \
        protocol/messages/event.proto \
        --micro_out=./protobuf \
        --gogo_out=M:./protobuf \
        --go-grpc_out=./protobuf \
        --go-grpc_out=./protobuf \
        --validate_out="lang=go:./protobuf" 


protoc \
        -I /usr/local/include \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/gogo/protobuf/gogoproto/ \
        -I protocol/accounts \
        -I protocol/messages \
        -I protocol/pipes \
        -I protocol/healthcheck \
        -I protocol/accounting \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/envoyproxy/protoc-gen-validate \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/protoc-gen-gotagger/proto \
        protocol/pipes/pipe.proto \
        --micro_out=./protobuf \
        --go-grpc_opt=require_unimplemented_servers=false \
        --go-grpc_out=./protobuf \
        --gogo_out=M:./protobuf \
        --validate_out="lang=go:./protobuf" 

protoc \
        -I /usr/local/include \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/gogo/protobuf/gogoproto/ \
        -I protocol/accounts \
        -I protocol/messages \
        -I protocol/pipes \
        -I protocol/healthcheck \
        -I protocol/accounting \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/envoyproxy/protoc-gen-validate \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/protoc-gen-gotagger/proto \
        protocol/healthcheck/healthcheck.proto \
        --micro_out=./protobuf \
        --go-grpc_opt=require_unimplemented_servers=false \
        --go-grpc_out=./protobuf \
        --gogo_out=M:./protobuf \
        --validate_out="lang=go:./protobuf"

protoc \
        -I /usr/local/include \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/gogo/protobuf/gogoproto/ \
        -I protocol/messages \
        -I protocol/pipes \
        -I protocol/healthcheck \
        -I protocol/accounts \
        -I protocol/accounting \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/envoyproxy/protoc-gen-validate \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/protoc-gen-gotagger/proto \
        protocol/accounts/user.proto \
        --micro_out=./protobuf \
        --go-grpc_opt=require_unimplemented_servers=false \
        --go-grpc_out=./protobuf \
        --gogo_out=M:./protobuf \
        --validate_out="lang=go:./protobuf"

protoc \
        -I /usr/local/include \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/gogo/protobuf/gogoproto/ \
        -I protocol/messages \
        -I protocol/pipes \
        -I protocol/healthcheck \
        -I protocol/accounts \
        -I protocol/accounting \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/envoyproxy/protoc-gen-validate \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/protoc-gen-gotagger/proto \
        protocol/accounts/apikey.proto \
        --micro_out=./protobuf \
        --go-grpc_opt=require_unimplemented_servers=false \
        --go-grpc_out=./protobuf \
        --gogo_out=M:./protobuf \
        --validate_out="lang=go:./protobuf" 

protoc \
        -I /usr/local/include \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/gogo/protobuf/gogoproto/ \
        -I protocol/messages \
        -I protocol/pipes \
        -I protocol/healthcheck \
        -I protocol/accounts \
        -I protocol/accounting \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/envoyproxy/protoc-gen-validate \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/protoc-gen-gotagger/proto \
        protocol/accounts/service.proto \
        --micro_out=./protobuf \
        --go-grpc_opt=require_unimplemented_servers=false \
        --go-grpc_out=./protobuf \
        --gogo_out=M:./protobuf \
        --validate_out="lang=go:./protobuf" 

protoc \
        -I /usr/local/include \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/gogo/protobuf/gogoproto/ \
        -I protocol/messages \
        -I protocol/pipes \
        -I protocol/healthcheck \
        -I protocol/accounts \
        -I protocol/accounting \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/envoyproxy/protoc-gen-validate \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/protoc-gen-gotagger/proto \
        protocol/accounting/accountingservice.proto \
        --micro_out=./protobuf \
        --go-grpc_opt=require_unimplemented_servers=false \
        --go-grpc_out=./protobuf \
        --gogo_out=M:./protobuf \
        --validate_out="lang=go:./protobuf" 

# protoc \
#         -I /usr/local/include \
#         -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/gogo/protobuf/gogoproto/ \
#         -I protocol/accounts \
#         -I protocol/messages \
#         -I protocol/pipes \
#         -I protocol/healthcheck \
#         -I protocol/accounting \
#         -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
#         -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/envoyproxy/protoc-gen-validate \
#         -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
#         -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/protoc-gen-gotagger/proto \
#         ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/gogo/protobuf/gogoproto/gogo.proto \
#         ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto/pmongo/objectid.proto \
#         protocol/accounting/accountingservice.proto \
#         protocol/accounts/apikey.proto \
#         protocol/accounts/service.proto \
#         protocol/accounts/user.proto \
#         protocol/healthcheck/healthcheck.proto \
#         protocol/messages/event.proto \
#         protocol/messages/messageenvelop.proto \
#         protocol/pipes/pipe.proto \
#         --js_out=import_style=commonjs,binary:./js/protobuf/ \
#         --grpc_out=generate_package_definition:.\output protocol/pipes/pipe.proto

grpc_tools_node_protoc \
        --js_out=import_style=commonjs,binary:./js/protobuf/ \
        --grpc_out=grpc_js:./js/protobuf \
        -I /usr/local/include \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/gogo/protobuf/gogoproto/ \
        -I protocol/accounts \
        -I protocol/messages \
        -I protocol/pipes \
        -I protocol/healthcheck \
        -I protocol/accounting \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/envoyproxy/protoc-gen-validate \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/protoc-gen-gotagger/proto \
        ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/gogo/protobuf/gogoproto/gogo.proto \
        ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto/pmongo/objectid.proto \
        protocol/accounting/accountingservice.proto \
        protocol/accounts/apikey.proto \
        protocol/accounts/service.proto \
        protocol/accounts/user.proto \
        protocol/healthcheck/healthcheck.proto \
        protocol/messages/event.proto \
        protocol/messages/messageenvelop.proto \
        protocol/pipes/pipe.proto 

mv protobuf/github.com/nochte/pipelinr-protocol/protobuf/* protobuf/
rm -rf protobuf/github.com

cp -rv extensions/* protobuf

# below this for mac / gvm
        # -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        # -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/envoyproxy/protoc-gen-validate \
        # -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        # -I ${HOME}/.gvm/pkgsets/go1.13/global/src/github.com/amsokol/protoc-gen-gotagger/proto \


#below for something else
        # -I ${HOME}/go/1.13.0/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        # -I ${HOME}/go/1.13.0/src/github.com/envoyproxy/protoc-gen-validate \
        # -I ${HOME}/go/1.13.0/src/github.com/amsokol/mongo-go-driver-protobuf/proto \
        # -I ${HOME}/go/1.13.0/src/github.com/amsokol/protoc-gen-gotagger/proto \
# -I ${HOME}/go/1.13.0/src/github.com/gogo/protobuf/gogoproto/ \