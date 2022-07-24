package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "heloo from home!!")
}

func main() {
	// chain := blockchain.GetBlockchain()
	// chain.AddBlock("Second Block")
	// chain.AddBlock("Third Block")

	// chain.AllBlocks()
	// for _, block := range chain.AllBlocks() {
	// 	fmt.Printf("%s\n", block.Data)
	// 	fmt.Printf("%s\n", block.Hash)
	// 	fmt.Printf("%s\n", block.PrevHash)
	// }

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
