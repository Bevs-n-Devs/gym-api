package handler

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	welcomeMsg := "Welcome to the Gym API!"
	w.Write([]byte(welcomeMsg))
}
