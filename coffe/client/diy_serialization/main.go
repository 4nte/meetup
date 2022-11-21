package main

import (
	"fmt"
	"log"
	"net"

	. "github.com/4nte/coffe/custom_serialization"
	"github.com/gookit/goutil/dump"
)

func main() {
	// Request for the coffe machine
	var makeCoffeRequest MakeCoffeRequest

	// Collect user input
	{
		fmt.Println("What type of coffe do you want? (cappuccino, long_black, americano)")
		fmt.Scanf("%s", &makeCoffeRequest.CoffeType)

		fmt.Println("How much sugar do you want in grams?")
		fmt.Scanf("%d", &makeCoffeRequest.AmountOfSugarInGrams)
	}

	dump.P(makeCoffeRequest)

	// Serialize the MakeCoffeRequest
	buffer := SerializeMakeCoffeeRequest(makeCoffeRequest)

	// Opens a new TCP connection with the coffe machine
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Panicf("failed to dial coffe machine: %s", err)
	}

	// Send the bits to the coffe machine
	// Eventually, these bits will travel over the wire as oscillations in voltage and/or electromagnetic waves over the air
	bytesWritten, err := conn.Write(buffer[:])
	if err != nil {
		log.Panicf("failed to send message: %s", err)
	}

	conn.Close()

	fmt.Printf("message sent (%d bytes)\n", bytesWritten)

}
