package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/joho/godotenv"
)

type fizzbuzzRequest struct {
	Count int `json:"count"`
}

type fizzbuzzResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	godotenv.Load()

	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz", fizzbuzzHandler).Methods("POST")

	handler := cors.Default().Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Server listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		panic(err)
	}
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	var req fizzbuzzRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Count <= 0 {
		http.Error(w, "count must be a positive integer", http.StatusBadRequest)
		return
	}

	fizzMsg := os.Getenv("FIZZ_MSG")
	buzzMsg := os.Getenv("BUZZ_MSG")
	fizzbuzzMsg := os.Getenv("FIZZBUZZ_MSG")

	var msg string
	if req.Count%3 == 0 && req.Count%5 == 0 {
		msg = fizzbuzzMsg
	} else if req.Count%3 == 0 {
		msg = fizzMsg
	} else if req.Count%5 == 0 {
		msg = buzzMsg
	}

	resp := fizzbuzzResponse{
		Status:  http.StatusOK,
		Message: msg,
	}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
