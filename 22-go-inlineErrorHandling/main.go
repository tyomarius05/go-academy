package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Inline Error Handling di Golang=====")

	fmt.Println()
	fmt.Println("Contoh Inline Error Handling bawaan Golang pada package strings")

	x := "10"
	if conv, err := strconv.Atoi(x); err != nil { //inline error handling
		fmt.Println("Terjadi kesalahan, ", err.Error())
	} else {
		fmt.Println(conv)
	}

	y := "b"
	if conv, err := strconv.Atoi(y); err != nil { //inline handling error
		fmt.Println("Terjadi kesalahan, ", err.Error())
	} else {
		fmt.Println(conv)
	}

	fmt.Println()
	fmt.Println("Contoh Error Handling yang kita bisa create")

	if result, err := div(10, 5); err != nil { //inline handling error
		fmt.Println("Terjadi kesalahan,", err.Error())
	} else {
		fmt.Println(result)
	}

	if result, err := div(10, 0); err != nil { //inline handling error
		fmt.Println("Terjadi kesalahan,", err.Error())
	} else {
		fmt.Println(result)
	}
}

func div(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divided by 0")
	}
	return (x / y), nil
}
