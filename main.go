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

const english = "English"
const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name string, language ...string) string {
	if name == "" {
		name = "World"
	}
	var l = english
	if len(language) == 1 {
		l = language[0]
	}
	return greetingPrefix(l) + name
}

func main() {
	fmt.Println(Hello("world"))
}
