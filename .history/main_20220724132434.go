package main

import (
	"fmt"
	"log"
	"net/http"
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
			URL: "/",
			MethodL "GET",
			DesDescription: "See Documentation",
		},
	}
	b, err := json.Marshal(data)
	utils.HandleErr(err)
}

func main() {
	http.HandleFunc("/")
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
