package main

import (
	"fmt"
	"log"
	"net"
)

func RunServer(address string, connHandler func(net.Conn)) {
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
		go connHandler(conn)
	}
}
