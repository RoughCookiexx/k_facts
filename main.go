package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type FactResponse struct {
	Message string `json:"message"`
}

func factHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request body", http.StatusBadRequest)
		return
	}
	bodyStr := string(bodyBytes)

	if strings.Contains(strings.ToLower(bodyStr), "roughc3potassium") {
		fact := GetRandomPotassiumFact()
		fmt.Printf("Fact: %s", fact)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(FactResponse{Message: fact})
		return
	}

	http.Error(w, "No potassium keyword found", http.StatusForbidden)
}

func sendSubscriptionRequest() {
	fmt.Println("Subscribing to chat message stream\n")
	otherAppSubscribeEndpoint := "http://127.0.0.1:6969/subscribe" // The other app's IP:PORT/path
	ourFactCallbackURL := "http://localhost:6970/fact"  // Our app's endpoint for receiving facts
                                                                  // Adjust host/port if this app isn't on localhost:8080
                                                                  // or if the other app needs a different address to reach this one.

	// Prepare the full URL with our callback as a query parameter
	data := url.Values{}
	data.Set("url", ourFactCallbackURL)
	requestURL := fmt.Sprintf("%s?%s", otherAppSubscribeEndpoint, data.Encode())

	// Send the POST request.
	// The other app's /subscribe endpoint (from previous context) expects 'url' in the query string.
	// Body is nil as data is in the URL.
	resp, err := http.Post(requestURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Error sending subscription request to %s: %v\n", requestURL, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Subscription request sent to %s. Status: %s\n", requestURL, resp.Status)
	// Optionally, read resp.Body for more details from the other app
}

func main() {
	fmt.Println("Starting k_facts\n")
	http.HandleFunc("/fact", factHandler)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	sendSubscriptionRequest()
	http.ListenAndServe(":6970", nil)
}

