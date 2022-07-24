package main

import "fmt"

func main() {
	foods := [3]string{"potato", "pizza", "pasta"}
	for _, food := range foods {
		fmt.Println(food)
	}
}
