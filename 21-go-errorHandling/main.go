package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Error Handling di Golang=====")

	fmt.Println()
	fmt.Println("Contoh Error Handling bawaan Golang pada package strings")

	x := "10"

	conv, err := strconv.Atoi(x)

	if err != nil {
		fmt.Println("Terjadi kesalahan, ", err.Error())
	} else {
		fmt.Println(conv)
	}

	y := "b"
	conv, err = strconv.Atoi(y)

	if err != nil {
		fmt.Println("Terjadi kesalahan, ", err.Error())
	} else {
		fmt.Println(conv)
	}

	fmt.Println()
	fmt.Println("Contoh Error Handling yang kita bisa create")
	result, err := div(10, 5)
	if err != nil {
		fmt.Println("Terjadi kesalahan,", err.Error())
	} else {
		fmt.Println(result)
	}

	result, err = div(10, 0)
	if err != nil {
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
