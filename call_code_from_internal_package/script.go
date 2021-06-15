package main

import (
	"fmt"
	"log"
	"github.com/accalina/golang-journey/create_go_module"
)

func main() {
	log.SetPrefix("Greeting: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Accalina")
	if err != nil { log.Fatal(err) }
	fmt.Println(message)
}