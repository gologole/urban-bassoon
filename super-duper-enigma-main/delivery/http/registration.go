package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const index = "index.html"

//запуск через горутину

func Handlereg() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, index)
	})

	http.HandleFunc("/handleClick", handleJSONRequest)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
	fmt.Println("Server is starting")
}

func handleJSONRequest(w http.ResponseWriter, r *http.Request) {
	// Check the request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON packet from the request body
	var data MyData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Process the JSON data
	fmt.Println("Received JSON data:")
	fmt.Println("Login:", data.Login)
	fmt.Println("Password:", data.Password)

	// Send a response
	response := map[string]interface{}{
		"message": "JSON received successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
