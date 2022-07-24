package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nomadcoders/nomadcoin/blockchain"
	"github.com/nomadcoders/nomadcoin/utils"
)

const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescirption struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description1"`
	Payload     string `json:"payload,omitempty"`
}

func (u URLDescirption) String() string {
	return "hello url description"
}

type AddBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescirption{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
	}
	fmt.Println(data)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody
		fmt.Println(addBlockBody)
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		fmt.Println(addBlockBody)
	}
}

func main() {
	fmt.Println(URLDescirption{
		URL:         "/",
		Method:      "GET",
		Description: "See Documentation",
	})
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
