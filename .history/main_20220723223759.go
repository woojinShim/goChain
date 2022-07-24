package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port string = ":4000"

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/home.html")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
