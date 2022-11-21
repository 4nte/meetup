package main

import (
	"log"
	"net"
)

/*
 _____        __  __                                 _     _
/  __ \      / _|/ _|                               | |   (_)
| /  \/ ___ | |_| |_ ___  ___   _ __ ___   __ _  ___| |__  _ _ __   ___
| |    / _ \|  _|  _/ _ \/ _ \ | '_ ` _ \ / _` |/ __| '_ \| | '_ \ / _ \
| \__/| (_) | | | ||  __|  __/ | | | | | | (_| | (__| | | | | | | |  __/
 \____/\___/|_| |_| \___|\___| |_| |_| |_|\__,_|\___|_| |_|_|_| |_|\___|
*/

func main() {
	// TCP server setup
	listen, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Panicf("failed to create the tcp server: %s", err)
	}

	// Accept all incoming TCP connections
	for {
		conn, err := listen.Accept() // This will block until a new connection arrives
		if err != nil {
			log.Panicf("failed to accept connection: %s", err)
		}

		// process the connection in a new go routine
		go handleIncomingConnection(conn)
	}
}

// This function is called for every connection
func handleIncomingConnection(conn net.Conn) {
	var buffer []byte

	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

}

// Types of coffe
type CoffeType int

const (
	CAPPUCCINO = 1
	LONG_BLACK = 2
	AMERICANO  = 3
)

type MakeCoffeRequest struct {
	CoffeType          string
	SugarAmountInGrams int8
}
