package main

import "fmt"

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	// genesisBlock := block{"Genesis Block", "", ""}
	// genesisBlock.hash = sha256.Sum256()
	for _, aByte := range "Genesis Block" {
		fmt.Println(aByte)
	}
}
