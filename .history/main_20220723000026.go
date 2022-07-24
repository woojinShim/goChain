package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main() {
	result := plus(2, 2)
	fmt.Println(result)
}
