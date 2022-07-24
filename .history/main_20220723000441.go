package main

import "fmt"

func plus(a ...int) int {
	return a + b
}

func main() {
	result := plus(2, 2, 4, 5, 6, 7, 8, 89)
	fmt.Println(result)
}
