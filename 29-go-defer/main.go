package main

import "fmt"

func main() {
	fmt.Println("=====Ini merupakan contoh Defer di Golang=====")
	fmt.Println()

	defer fmt.Println("Ini akan dieksekusi terakhir ke 2")

	defer fmt.Println("Ini akan dieksekusi terakhir ke 1")

	fmt.Println("Cetak tanpa defer")

	/*
		Konsep dari defer ini LIFO (Last In First Out), jadi defer yang paling atas akan dieksekusi paling terakhir
	*/
}
