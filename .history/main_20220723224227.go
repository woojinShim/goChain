package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/home.html")
	if err != nil {
		log.Fatal(err)
	}
	data := homeData{"Home"}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
