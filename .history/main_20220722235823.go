package main

import "fmt"

func plus(a, b int, name string) (int, string) {
	return a + b, name
}

func main() {
	result := plus(2, 2, name)
	fmt.Println(result)
}
