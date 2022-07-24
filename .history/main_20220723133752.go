package main

import "fmt"

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	genesisBlock := block{"Genesis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))

	}
}
