package main

type block struct {
	data    string
	hash    string
	preHash string
}

func main() {
	genesisBlock := block{"Genesis Block", "", ""}
	genesisBlock.hash = fn(genesisBlock.data + genesisBlock.preHash)
}
