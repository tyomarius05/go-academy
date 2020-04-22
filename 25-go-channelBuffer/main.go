package main

import "fmt"

func main() {
	fmt.Println("=====Ini merupakan contoh Goroutine with Channel di Golang=====")
	fmt.Println()

	//contoh channel buffer
	myChan := make(chan string, 2)
	myChan <- "Hallo"
	myChan <- "Golang"
	close(myChan) //ini untuk close channel

	fmt.Println(<-myChan)
	fmt.Println(<-myChan)

	//jika diprint lagi maka tidak akan muncul datanya, karena sudah kluar datanya/diprint
	//jadi channel itu jika sudah dikeluarkan maka tidak bisa digunakan lagi valuenya.
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
}
