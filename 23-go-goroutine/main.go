package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Goroutine di Golang=====")
	fmt.Println()

	go helloGo()

	time.Sleep(1 * time.Second) //fungsi untuk menunggu hasil balikan dari goroutine helloGo(), jika tidak menggunkan ini tidak tampil hasil dari Goroutinenya
	//dan sebenarnya ini bukan best practice.
	fmt.Println("Dicetak after Goroutine go helloGo()")
}

func helloGo() {
	fmt.Println("Hallo Tyo Marius")
}
