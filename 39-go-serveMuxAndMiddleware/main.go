package main

import (
	"fmt"
	"log"
	"net/http"
)

//Contoh middleware, disini kita memisalkan user hit suatu endpoint sebenernya masuk ke suatu endpoint untuk mencatat log, kemudian baru masuk
//ke endpoint sebenernya
func logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("%s request %s", req.RemoteAddr, req.URL)
		h.ServeHTTP(res, req)
	})
}

func main() {
	fmt.Println("=====Ini merupakan contoh HTTP Server with ServeMux and Middleware di Golang=====")
	fmt.Println()

	h := http.NewServeMux() //sebenernya ini juga http.Handler (interface) karena ini polimorfismenya http.Handler

	h.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Sekarang kita menggunakan ServerMux")
	})

	h.HandleFunc("/test1", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "EndPoint Test 1")
	})

	h.HandleFunc("/test2", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "EndPoint Test 2")
	})

	h.HandleFunc("/test3", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "EndPoint Test 3")
	})

	h1 := logRequest(h)

	log.Println("Running server 3000")

	err := http.ListenAndServe(":3000", h1)

	if err != nil {
		fmt.Println(err.Error())
	}
}
