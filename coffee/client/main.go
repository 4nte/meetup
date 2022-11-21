package main

import (
	"fmt"

	. "github.com/4nte/meetup/coffee/diy_serialization"
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

	SendMessageToCoffeeMachine(buffer[:])

}
