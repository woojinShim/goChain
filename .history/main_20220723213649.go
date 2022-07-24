package main

import (
	"fmt"

	"github.com/nomadcoders/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	fmt.Println(chain)
}
