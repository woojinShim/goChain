package main

import (
	"fmt"

	"github.com/nomadcoders/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	fmt.Println(chain)
}
