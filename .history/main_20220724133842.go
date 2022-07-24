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
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(data)
	utils.HandleErr(err)
	fmt.Fprintf(rw, "%s", b)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
