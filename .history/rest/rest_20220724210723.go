package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nomadcoders/nomadcoin/blockchain"
	"github.com/nomadcoders/nomadcoin/utils"
)

const port string = ":4000"

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescirption struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description1"`
	Payload     string `json:"payload,omitempty"`
}

func (u urlDescirption) String() string {
	return "hello url description"
}

type addBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescirption{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{id}"),
			Method:      "GET",
			Description: "See A Block",
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
		var addBlockBody addBlockBody

		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
		fmt.Println(addBlockBody)
	}
}

func Start(aPort int) {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
