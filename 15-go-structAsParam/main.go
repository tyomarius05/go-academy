package main

import (
	"fmt"
)

type player struct {
	ID           int
	PlayerName   string
	Position     string
	OverallPoint float32
}

func main() {
	fmt.Println("=====Ini merupakan contoh Struct as Parameter/Argument di Golang=====")

	p := player{
		ID:           1,
		PlayerName:   "Tyo Marius",
		Position:     "CF",
		OverallPoint: 99.9,
	}

	printPlayer(p)
}

func printPlayer(p player) {
	fmt.Println("Player ID =", p.ID)
	fmt.Println("Player Name =", p.PlayerName)
	fmt.Println("Player Position =", p.Position)
	fmt.Println("Player Overall Point =", p.OverallPoint)
}
