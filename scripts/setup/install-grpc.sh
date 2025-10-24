PROTOC_ZIP=protoc-25.2-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v25.2/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr 'include/*'
rm -f $PROTOC_ZIP

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5