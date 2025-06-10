package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RoughCookiexx/twitch_chat_subscriber"
)

type FactResponse struct {
	Message string `json:"message"`
}

func sendSubscriptionRequest() {
	fmt.Println("Subscribing to chat message stream\n")
	targetURL := "http://localhost:6969/subscribe"
	filterPattern := "Potassium"
	
	response, err := twitch_chat_subscriber.SendRequestWithCallbackAndRegex(targetURL, GetRandomPotassiumFact, filterPattern, 6974)
	
	if err != nil {
		fmt.Errorf("Failed to subscribe to Twitch chat message stream.\n%s", err)
	}

	fmt.Printf("Subscribed to Twitch chat message stream. Status: %s Err: %s\n", response, err)
}

func main() {
	fmt.Println("Starting k_facts\n")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	sendSubscriptionRequest()
	http.ListenAndServe(":6974", nil)
}

