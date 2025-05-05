package handler

import (
	"log"
	"net/http"
	"os"
)

const (
	localPort = "9009"
)

func StartServer() {
	log.Println("starting server")
	// define routes
	log.Println("defining routes")
	http.HandleFunc("/", Home)
	http.HandleFunc("/gymapi", GymAPI)
	http.HandleFunc("/gymapi/ping", Ping)
	http.HandleFunc("/gymapi/check", Check)

	// initialise port for application
	httpPort := os.Getenv("PORT")

	// start server on machine if hosting platform port from hosting platform
	if httpPort == "" {
		log.Println("No PORT environment variable detected, using local port")
		httpPort = localPort
	}
	log.Println("HTTP server running on http://localhost:" + httpPort)

	// start server on hosting platform port
	err := http.ListenAndServe(":"+httpPort, nil)
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err.Error())
	}
}
