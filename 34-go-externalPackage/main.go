package main

import (
	"crypto/sha1"
	"fmt"

	p "github.com/wuriyanto48/go-pbkdf2"
)

func main() {
	fmt.Println("=====Ini merupakan contoh External Package di Golang=====")
	fmt.Println()

	pass := p.NewPassword(sha1.New, 8, 32, 15000)
	hashed := pass.HashPassword("tyomariusGoAcademy")
	fmt.Println(hashed.CipherText)
	fmt.Println(hashed.Salt)
}
