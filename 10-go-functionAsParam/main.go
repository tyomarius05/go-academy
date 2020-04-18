package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Function As Arguments/Parameter di Golang=====")

	myFunc := func(kata string) bool {
		return kata == "Golang"
	}

	result := match("Golang", myFunc)

	if result {
		fmt.Println("Matched")
	} else {
		fmt.Println("Not Matched")
	}
}

func match(kata string, f func(string) bool) bool {
	if f(kata) {
		return true
	}
	return false
}
