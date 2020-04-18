package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=====Ini merupakan contoh String Manipulation di Golang=====")

	var x string = "Tyo Marius"

	if strings.EqualFold(x, "tyo marius") { //EqualFold digunakan untuk membandingkan dua string dengan tidak case sensitive
		fmt.Println("Sama")
	} else {
		fmt.Println("Tidak sama")
	}

	fmt.Println(strings.ToLower(x))
	fmt.Println(strings.ToUpper(x))

	if strings.Contains(x, "golang") {
		fmt.Println("Kata", x, "mengandung kata golang")
	} else {
		fmt.Println("Kata", x, "tidak mengandung kata golang")
	}
}
