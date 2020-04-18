package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Function As Value di Golang=====")

	//Ini contoh basic function yang telah kita pelajari
	a := 6
	b := 4
	fmt.Println(swapBasicFunc(a, b))

	//Ini contoh Function As Value
	swapFunc := func(a, b int) (int, int) {
		return b, a
	}

	fmt.Println(swapFunc(a, b))

	hello := func(kata string) string {
		return fmt.Sprintf("Hello %s", kata)
	}

	fmt.Println(hello("Tyo Marius"))
}

func swapBasicFunc(a, b int) (int, int) {
	return b, a
}
