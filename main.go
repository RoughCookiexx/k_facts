package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type FactResponse struct {
	Fact string `json:"fact"`
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(FactResponse{Fact: fact})
		return
	}

	http.Error(w, "No potassium keyword found", http.StatusForbidden)
}

func main() {
	http.HandleFunc("/fact", factHandler)
	http.ListenAndServe(":8080", nil)
}

