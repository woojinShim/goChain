package main

import (
	"fmt"

	"github.com/nomadcoders/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AllBlocks()
	for _, block := range chain.AllBlocks() {
		fmt.Printf("%s", block.Data)
		fmt.Printf("%s", block.Hash)
		fmt.Printf("%s", block.PrevHash)
	}
}
