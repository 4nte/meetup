rm -rf ./gen
buf generate proto/coffee/$1
cd ./gen/go/github.com/4nte/meetup/coffee/proto
go mod init github.com/4nte/meetup/coffee/proto
go mod tidy

echo "Sucessfully generated proto files"