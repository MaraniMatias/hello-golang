package main

import "fmt"

type person struct {
	name string
	age  int
}

// p := Person("Matias", 26)
func Person(name string, age int) *person {
	p := person{name: name}
	p.age = age
	return &p
}

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
