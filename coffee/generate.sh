rm -rf ./gen
buf generate proto/coffe/v1
cd ./gen/go/github.com/4nte/coffe/proto
go mod init github.com/4nte/coffe/proto
go mod tidy
