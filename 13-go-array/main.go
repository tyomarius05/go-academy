package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Array di Golang=====")

	var arrInt [5]int
	arrInt[0] = 1
	arrInt[1] = 2
	arrInt[2] = 3
	arrInt[3] = 4
	arrInt[4] = 5

	for i, v := range arrInt {
		fmt.Println("Index ke-", i, "=", v)
	}

	arrString := [3]string{"Tyo", "Marius", "Golang"}

	for _, v := range arrString {
		fmt.Println(v)
	}

	arrMulti := [2][2]int{{1, 2}, {3, 4}}
	for _, v := range arrMulti {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	for i, v := range arrMulti {
		for i2, v2 := range v {
			fmt.Println("Indeks ", i, i2, "=", v2)
		}
	}
}
