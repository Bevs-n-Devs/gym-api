package handler

import (
	"encoding/json"
	"net/http"
)

// map object to mock database
var blacklistedUsers = map[string]bool{
	"101": true,
	"103": true,
	"105": true,
}

var notBlacklistedUsers = map[string]bool{
	"102": false,
	"104": false,
	"106": false,
}

type CheckResponse struct {
	Banned     bool   `json:"banned"`
	Membership string `json:"membership"`
	Status     int    `json:"status"`
}

func Check(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userid := r.URL.Query().Get("userid")
	if userid == "" {
		http.Error(w, "Missing userid parameter", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if userid != "" {
		userBanned := validateUserid(userid)

		if userBanned {
			response := CheckResponse{
				Banned:     true,
				Membership: "blacklisted",
				Status:     http.StatusOK,
			}

			// marshal the response to JSON
			jsonResonse, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
				return
			}
			// write the JSON response
			w.Write(jsonResonse)
		}

		if !userBanned {
			response := CheckResponse{
				Banned:     false,
				Membership: "not blacklisted",
				Status:     http.StatusOK,
			}

			// marshal the response to JSON
			jsonResonse, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
				return
			}
			// write the JSON response
			w.Write(jsonResonse)
		}
	}
}

// returns true if userid is valid
func validateUserid(userid string) bool {
	var result bool
	if blacklistedUsers[userid] {
		result = true
	}

	if notBlacklistedUsers[userid] {
		result = false
	}

	return result
}
