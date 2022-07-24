package main

import (
	"fmt"

	"github.com/nomadcoders/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	chain.AllBlocks()
	for _, block := range chain.AllBlocks() {
		fmt.Printf("%s\n", block.Data)
		fmt.Printf("%s\n", block.Hash)
		fmt.Printf("%s\n", block.PrevHash)
	}
}
