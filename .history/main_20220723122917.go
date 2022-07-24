package main

import (
	"fmt"

	"github.com/nomadcoders/nomadcoin/person"
)

func main() {
	nico := person.Person{"nico", 12}
	fmt.Println(nico)
}
