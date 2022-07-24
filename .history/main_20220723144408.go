package main

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.block[len(b.blocks)-1].hash
	}
}

func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", ""}
	if len(b.blocks) > 0 {
		newBlock.prevHash = b.blocks[len(b.blocks)-1].hash
	}
}

func (b *blockchain) listBlocks() {

}

func main() {
	chain := blockchain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
}
