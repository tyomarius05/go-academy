package main

import "fmt"

func main() {
	fmt.Println("===== Ini merupakan latihan Unit Test Ke-2 dengan Golang =====")
	fmt.Println()

	fmt.Println(add(2, 3))
}

func add(x, y int) int {
	return x + y
}
