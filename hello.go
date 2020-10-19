package main

import "fmt"

const englishHelloPrefix = "Hello, "

/*
Hello World printing
*/
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
