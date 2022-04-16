package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
func main() {
	//Greet(os.Stdout, "Sinisa")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}

func Greet(writer io.Writer, name string) {
	//fmt.Printf("Hello, %s!", name)
	fmt.Fprintf(writer, "Hello, %s!", name)
}
