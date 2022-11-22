module github.com/4nte/meetup/coffee/client

go 1.19

require (
	github.com/4nte/meetup/coffee/diy_serialization v0.0.0-00010101000000-000000000000
	github.com/4nte/meetup/coffee/proto v0.0.0-00010101000000-000000000000
	github.com/gookit/goutil v0.5.15
	google.golang.org/protobuf v1.28.1
)

replace github.com/4nte/meetup/coffee/diy_serialization => ../diy_serialization

replace github.com/4nte/meetup/coffee/proto => ../gen/go/github.com/4nte/meetup/coffee/proto

require (
	github.com/gookit/color v1.5.2 // indirect
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261 // indirect
	golang.org/x/text v0.3.8 // indirect
)
