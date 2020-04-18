package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Basic Function di Golang=====")

	x := 6
	y := 4

	result := add(x, y)
	fmt.Println(result)

	fmt.Println(mod(x, y))

	fmt.Println("Contoh Function yang bisa mereturnkan 2 value")
	kata1 := "Tyo"
	kata2 := "Golang"

	fmt.Println(kata1, kata2)
	fmt.Println(swap(kata1, kata2))

	resultKata1, resultKata2 := swap(kata1, kata2)
	fmt.Println(resultKata1, resultKata2)
}

func add(a int, b int) int {
	return a + b
}

func mod(a, b int) int {
	return a % b
}

func swap(kata1, kata2 string) (string, string) {
	return kata2, kata1
}
