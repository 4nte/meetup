package main

/*
 _____        __  __                                 _     _
/  __ \      / _|/ _|                               | |   (_)
| /  \/ ___ | |_| |_ ___  ___   _ __ ___   __ _  ___| |__  _ _ __   ___
| |    / _ \|  _|  _/ _ \/ _ \ | '_ ` _ \ / _` |/ __| '_ \| | '_ \ / _ \
| \__/| (_) | | | ||  __|  __/ | | | | | | (_| | (__| | | | | | | |  __/
 \____/\___/|_| |_| \___|\___| |_| |_| |_|\__,_|\___|_| |_|_|_| |_|\___|
*/

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"

	diy "github.com/4nte/meetup/coffee/diy_serialization"
	coffe "github.com/4nte/meetup/coffee/proto"
	"github.com/4nte/meetup/coffee/server/api"
	protobuf "google.golang.org/protobuf/proto"
)

func main() {
	//RunServer("localhost:9000", diyConnectionHandler)
	//RunServer("localhost:9000", protobufConnectionHandler)
	api.Setup()

}

// This function is called for every connection
func diyConnectionHandler(conn net.Conn) {
	r := bufio.NewReader(conn)

	for {
		messageIdentifier, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panicf("error reading tag byte: %s", err)
		}
		fmt.Printf("first byte is %b\n", messageIdentifier)

		var buffer = make([]byte, diy.MAKE_COFFE_REQUEST_PAYLOAD_SIZE)

		// 1. Wait for the message to arrive
		// 2. Copy the message to the `buffer` variable
		_, err = io.ReadFull(r, buffer)

		switch err {
		case nil:
		case io.EOF:
			return
		default:
			log.Panicf("failed to read from the socket: %s", err)
		}

		// Print to stdout (for debugging)
		fmt.Printf("Message received, size: %d, data: %b\n", len(buffer), buffer)

		request := diy.DeserializeMakeCoffeeRequest(buffer)
		fmt.Println(request)
	}
}

func protobufConnectionHandler(conn net.Conn) {
	r := bufio.NewReader(conn)

	// Read data from the socket
	data, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	// Print the received bytes
	fmt.Printf("Message received, size: %d, data: %b\n", len(data), data)

	var request coffe.MakeCoffeRequest
	err = protobuf.Unmarshal(data, &request)
	if err != nil {
		log.Fatalf("faield to deserialize: %s", err)
	}

	fmt.Printf("Deserialized message: {%v}\n\n", &request)

}
