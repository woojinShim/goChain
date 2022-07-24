package blockchain

import "sync"

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

var b *blockchain
var once sync.Once

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
		})

	}
	return b
}
