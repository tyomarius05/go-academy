package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=====Ini merupakan contoh External Package di Golang=====")
	fmt.Println()

	username := flag.String("U", "", "Your Username")
	password := flag.String("P", "", "Your Password")

	flag.Parse()

	resultLogin := login(*username, *password)

	if resultLogin {
		fmt.Println("Sukses login")
	} else {
		fmt.Println("Gagal login")
	}
}

func login(username, password string) bool {
	if strings.EqualFold(username, "tyomarius05") && strings.EqualFold(password, "12345") {
		return true
	}
	return false
}
