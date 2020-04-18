package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Looping di Golang=====")

	fmt.Println("A. Contoh case FooBar, jika angka kelipatan 3 cetak Foo, jika angka kelipaatan 5 cetak Bar, jika kelipatan 3 dan 5 cetak FooBar")
	for i := 1; i <= 30; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "= FooBar")
		} else if i%3 == 0 {
			fmt.Println(i, "= Foo")
		} else if i%5 == 0 {
			fmt.Println(i, "= Bar")
		}
	}

	fmt.Println("\nB. Contoh kedua")
	y := 0
	for y < 5 {
		fmt.Println(y)
		y++
	}

	fmt.Println("\nC. Contoh ketiga, penggunaan for seperti foreach di golang")
	contohArray := [5]string{
		"Pertama", "Kedua", "Ketiga", "Keempat", "Kelima",
	}

	fmt.Println("Jumlah data pada Array yaitu", len(contohArray))

	for index, value := range contohArray { //Jika membutuhkan nilai index
		fmt.Println("Nilai index", index, "yaitu", value)
	}

	fmt.Println()

	for _, value := range contohArray { //Jika membutuhkan tidak nilai index, maka balikan Index bisa kita ganti dengan _
		fmt.Println(value)
	}
}
