rm -rf ./gen
buf generate proto/coffee/v1
cd ./gen/go/github.com/4nte/coffee/proto
go mod init github.com/4nte/coffee/proto
go mod tidy
