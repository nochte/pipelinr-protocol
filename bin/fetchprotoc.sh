# PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
PROTOC_ZIP=protoc-3.15.8-linux-x86_64.zip
curl -OL https://github.com/google/protobuf/releases/download/v3.15.8/$PROTOC_ZIP
# unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
# unzip -o $PROTOC_ZIP -d /usr/local include/*
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local include/*
rm -f $PROTOC_ZIP