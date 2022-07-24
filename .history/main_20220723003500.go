package main

import "fmt"

func plus(a ...int) int {
	var total int
	for _, item := range a {
		total += item
	}
	return total
}

func main() {
	name := "asdf;lkjasdf;lkjasdf"
	for _, letter := range name {
		fmt.Println(string(letter))
	}
}
