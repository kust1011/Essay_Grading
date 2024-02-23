package main

import (
	"encoding/json"
	"net/http"
	"fmt"

	// "github.com/go-zeromq/zmq4/security/null"
)

type ThreadIDResponse struct {
    ThreadID string `json:"threadID"`
}

type UserInputRequest struct {
    UserInput string `json:"userInput"`
    ThreadID  string `json:"threadID"`
}

func getThreadID(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

	thread, err := CreateThread()
	if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Thread ID: %s\n", thread.ID)
	
	response := ThreadIDResponse{
		ThreadID: thread.ID,
	}

	// Serialize response structure to JSON
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Error marshalling JSON: %v", err)
        return
    }

    // Set the response header to JSON and send the response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}


func handleThreadMessagesRequest(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
	
	var request UserInputRequest

	// Parse the request body into a structure
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	// Handle UserInput as a string
    userInput := request.UserInput
    threadID := request.ThreadID
    fmt.Println("Received user input:", userInput)
    fmt.Println("Received thread ID:", threadID)

	response, err := SendMessages(threadID, userInput)
	if err != nil {
		fmt.Println("Error:", err)
	}

    // Set the response header to JSON and send the response
    w.Header().Set("Content-Type", "application/json")
	
	jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }else{
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func main() {
    http.HandleFunc("/api/get-thread-id", getThreadID)
    http.HandleFunc("/api/thread-messages", handleThreadMessagesRequest)
	http.ListenAndServe(":8080", nil)
}
