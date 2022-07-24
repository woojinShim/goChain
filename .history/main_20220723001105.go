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
	// result := plus(2, 2, 4, 5, 6, 7, 8, 89)
	// fmt.Println(result)
	name := "asdf;lkjasdf;lkjasdf"
	for index, letter := range name {
		fmt.Println(string(letter))
	}
}
