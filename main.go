package main

import (
	"fmt"
	"log"

	"github.com/mjmhtjain/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, error := greetings.Hello("Mohit")

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
