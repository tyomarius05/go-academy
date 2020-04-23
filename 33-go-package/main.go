package main

import (
	"fmt"
	postgressConn "github.com/tyomarius05/go-academy/33-go-package/config"
	mPlayer "github.com/tyomarius05/go-academy/33-go-package/model"
)
func main(){
	fmt.Println("=====Ini merupakan contoh Custom Package di Golang=====")
	fmt.Println()

	conn := postgressConn.GetPostgresConnection()
	fmt.Println(conn)

	
	myPlayer := mPlayer.Player{ID:1, Name: "Tyo Marius"}

	fmt.Println(myPlayer)
}