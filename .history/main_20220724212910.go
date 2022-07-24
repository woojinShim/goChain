package main

import (
	"github.com/nomadcoders/nomadcoin/explorer"
	"github.com/nomadcoders/nomadcoin/rest"
)

func main() {
	explorer.Start(4000)
	rest.Start(5000)
}
