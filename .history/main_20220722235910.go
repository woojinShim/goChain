package main

import "fmt"

func plus(a, b int, name string) (int, string) {
	return a + b, name
}

func main() {
	result, name := plus(2, 2, "leiws")
	fmt.Println(result, name)
}
