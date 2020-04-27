package main

import (
	"fmt"
	"log"
	"net/http"
)

type myHandler struct{}

func (m myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello Golang")
}

func main() {
	fmt.Println("=====Ini merupakan contoh HTTP Server di Golang=====")
	fmt.Println()

	h := myHandler{}

	err := http.ListenAndServe(":3000", h)
	log.Fatal(err)
}
