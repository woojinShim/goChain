package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nomadcoders/nomadcoin/utils"
)

const port string = ":4000"

type URLDescirption struct {
	URL         string
	Method      string
	Description string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescirption{
		{
			URL:            "/",
			MethodL:        "GET",
			DesDescription: "See Documentation",
		},
	}
	b, err := json.Marshal(data)
	utils.HandleErr(err)
	fmt.Print(b)
}

func main() {
	http.HandleFunc("/")
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}