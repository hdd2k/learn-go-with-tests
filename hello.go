package main

import (
	"fmt"
)

const englishPrefixHello = "Hello, "

func Hello(name string) string {
	return englishPrefixHello + name
}

func main() {
	fmt.Println(Hello("world"))
}
