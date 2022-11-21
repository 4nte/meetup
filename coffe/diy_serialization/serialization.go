package custom_serialization

import "log"

const (
	MAKE_COFFE_REQUEST_IDENTIFIER = 1

	MAKE_COFFE_REQUEST_PAYLOAD_SIZE = 2
)

func SerializeMakeCoffeeRequest(request MakeCoffeRequest) [3]byte {
	var buffer [3]byte

	// Set the message identifier
	buffer[0] = 0b00000001

	// 1. Encode CoffeType in the first byte of the message
	switch request.CoffeType {
	case "cappuccino":
		buffer[1] = 0b00000001
	case "long_black":
		buffer[1] = 0b00000010
	case "americano":
		buffer[1] = 0b00000011
	default:
		log.Panicf("unknown coffe type %s", request.CoffeType)
	}

	// 2. Write AmountOfSugarInGrams in the second byte
	buffer[2] = byte(request.AmountOfSugarInGrams)

	return buffer
}

func DeserializeMakeCoffeeRequest(buffer []byte) MakeCoffeRequest {
	var makeCoffeRequest MakeCoffeRequest

	// 1. Decode CoffeType from the first byte
	switch buffer[1] {
	case 0b00000001:
		makeCoffeRequest.CoffeType = "cappuccino"
	case 0b00000010:
		makeCoffeRequest.CoffeType = "long_black"
	case 0b00000011:
		makeCoffeRequest.CoffeType = "americano"
	default:
		log.Panicf("unknown coffe type %b", buffer[0])
	}

	// 2. Decode AmountOfSugarInGrams from the second byte
	makeCoffeRequest.AmountOfSugarInGrams = int8(buffer[1])

	return makeCoffeRequest
}
