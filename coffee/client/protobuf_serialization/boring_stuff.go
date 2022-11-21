	package main
	

	func SendMessageToCoffeMachine([]byte) {

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
	}