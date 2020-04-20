package main

import (
	"fmt"
)

type vector struct {
	X int
	Y int
}

type player struct {
	ID           int
	PlayerName   string
	Position     string
	OverallPoint float32
}

func main() {
	fmt.Println("=====Ini merupakan contoh Struct di Golang=====")

	var v vector
	v.X = 2
	v.Y = 3

	fmt.Println(v.X, v.Y)

	p := player{ID: 1, PlayerName: "Tyo Marius", Position: "CF", OverallPoint: 99.9}

	fmt.Println("Player ID =", p.ID)
	fmt.Println("Player Name =", p.PlayerName)
	fmt.Println("Player Position =", p.Position)
	fmt.Println("Player Overall Point =", p.OverallPoint)
}
