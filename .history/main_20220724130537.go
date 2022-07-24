package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func main() {
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
