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

	. "github.com/4nte/coffe/custom_serialization"
)

func main() {
	RunServer("localhost:9000")
}

// This function is called for every connection
func HandleConnection(conn net.Conn) {
	r := bufio.NewReader(conn)

	for {
		messageIdentifier, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panicf("error reading tag byte: %s", err)
		}

		var payloadLength int
		{
			switch messageIdentifier {
			case 0b00000001:
				payloadLength = MAKE_COFFE_REQUEST_PAYLOAD_SIZE
			case 0b00000010:
				payloadLength = MAKE_COFFE_REQUEST_PAYLOAD_SIZE
			}

		}

		var buffer = make([]byte, payloadLength)

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
		fmt.Printf(`
		Message received:
		- size: %d bytes
		- buffer: %b
		`, len(buffer), buffer)

		request := DeserializeMakeCoffeeRequest(buffer)
		fmt.Println(request)

	}

}

func RunServer(address string) {
	// TCP server setup
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Panicf("failed to create the tcp server: %s", err)
	}
	fmt.Println("Coffe machine is ready to accept connections")

	// Accept all incoming TCP connections
	for {
		conn, err := listen.Accept() // This will block until a new connection arrives
		if err != nil {
			log.Panicf("failed to accept connection: %s", err)
		}

		// process the connection in a new go routine
		go HandleConnection(conn)
	}
}
