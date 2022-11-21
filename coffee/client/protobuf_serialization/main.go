package main

import "github.com/4nte/coffee/proto"

func main() {

	makeCoffeRequest := proto.MakeCoffeRequest{
		Type: proto.CoffeType_COFFE_TYPE_AMERICANO,
	}

	SendMessageToCoffeeMachine()
}
