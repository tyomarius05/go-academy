package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("=====Ini merupakan contoh HTTP Server with HandlerFunc di Golang=====")
	fmt.Println()

	h := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hallo Tyo Marius...")
	})

	log.Println("Server running on port 3000")

	err := http.ListenAndServe(":3000", h)

	if err != nil {
		log.Fatal(err)
	}
}
