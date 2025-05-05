package main

import (
	"log"

	"github.com/Bevs-n-Devs/gym-api/handler"
)

func main() {
	log.Println("Starting Gym API...")
	go handler.StartServer()

	// infinite loop to keep the server running
	select {}
}
