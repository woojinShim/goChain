package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescirption struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description1"`
	Payload     string `json:"payload,omitempty"`
}

func (u URLDescirption) String() string {
	return "hello url description"
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescirption{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/blocks",
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
