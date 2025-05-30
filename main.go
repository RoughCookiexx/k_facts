package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/RoughCookiexx/twitch_chat_subscriber"
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
	targetURL := "http://localhost:6969/subscribe"
	callbackURL := "http://localhost:6970/fact"
	filterPattern := "Potassium"
	
	response, err := twitch_chat_subscriber.SendRequestWithCallbackAndRegex(targetURL, callbackURL, filterPattern)
	
	if err != nil {
		fmt.Errorf("Failed to subscribe to Twitch chat message stream.\n%s", err)
	}

	fmt.Printf("Subscribed to Twitch chat message stream. Status: %s Err: %s\n", response)
}

func main() {
	fmt.Println("Starting k_facts\n")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.HandleFunc("/fact", factHandler)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	sendSubscriptionRequest()
	http.ListenAndServe(":6970", nil)
}

