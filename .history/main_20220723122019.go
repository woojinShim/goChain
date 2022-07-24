package main

import "fmt"

type person struct {
	name string
	age  int
}

func (inst person) sayHello() {
	fmt.Printf("Hello! My name is %s and I'm %d", inst.name, inst.age)
}

func main() {
	nico := person{"nico", 12}
	nico.sayHello()
}
