package main

import (
	"fmt"
	"io"
	// "log"
	"net/http"
	// "os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "hello world")
}

// func main() {
// 	// Greet(os.Stdout, "Hank2")
// 	log.Fatal(http.ListenAndServe(":5005", http.HandlerFunc(MyGreetHandler)))
// }
