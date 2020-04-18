package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Operator Aritmatika di Golang=====")

	a := 100
	b := 400
	result := a + b
	fmt.Println(result) //ini contoh jika menggunakan variable penampung
	fmt.Println(a + b)  //ini contoh jika tidak menggunakan variable penampung

	kata1 := "Hallo"
	kata2 := "Tyo Marius"
	fmt.Println(kata1 + " " + kata2)

	result = b - a
	fmt.Println(result)

	result = b / a
	fmt.Println(result)

	result = b % a
	fmt.Println(result)

	a = a + 1
	fmt.Println(a)

	a++
	fmt.Println(a)

	fmt.Println(a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}
