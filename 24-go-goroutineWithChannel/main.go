package main

import "fmt"

func main() {
	fmt.Println("=====Ini merupakan contoh Goroutine with Channel di Golang=====")
	fmt.Println()

	//ini contoh channel unbuffer
	done := make(chan bool) //channel digunakan agar goroutine main dan goroutine helloGo bisa tersynchronize

	go helloGo(done)
	<-done //jadi jika goroutine helloGo sudah selesai maka akan memberitahukan ke Channel bahwa sudah selesai dan goroutine Main bisa melanjutkan

	fmt.Println("Dicetak after menjalankan Goroutine")
}

func helloGo(done chan bool) {
	fmt.Println("Hello Tyo Marius")
	done <- true //memberitahukan ke channel bahwa goroutine selesai
}
