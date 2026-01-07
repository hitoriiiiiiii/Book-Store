package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if r.Body == nil {
		log.Println("Warning: empty request body")
		return
	}
	defer r.Body.Close() // important

	if err := json.NewDecoder(r.Body).Decode(x); err != nil {
		log.Printf("Failed to parse JSON body: %v\n", err)
	}
}
