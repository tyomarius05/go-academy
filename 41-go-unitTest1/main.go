package main

import (
	"fmt"
)

func main() {
	fmt.Println("===== Ini merupakan latihan membuat Unit Test dengan Golang =====")
	fmt.Println()

	result := calculate(5)
	fmt.Println(result)
}

func calculate(x int) int {
	return x + 2
}
