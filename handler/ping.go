package handler

import "net/http"

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	message := "GymAPI is running"
	w.Write([]byte(message))
}
