package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Function Closure di Golang=====")

	nextValue := genNumber() //variable Next Value disini bertipe Func, karena func genNumber mereturnkan suatu Func
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())

	fmt.Println()

	z := love("Tyo")

	fmt.Println(z("Intan"))
	fmt.Println(z("Golang"))
}

func genNumber() func() int {
	x := 0
	return func() int {
		x = x + 1
		return x
	}
}

func love(kata1 string) func(string) string {
	return func(kata2 string) string {
		return fmt.Sprintf("%s love %s", kata1, kata2)
	}
}
