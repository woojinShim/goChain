package main

import (
	"fmt"
	"net/http"
)

const port string = ":4000"

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
	http.ListenAndServe(port, nil)
}
