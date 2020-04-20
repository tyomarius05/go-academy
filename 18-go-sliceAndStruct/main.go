package main

import (
	"fmt"
	"strings"
)

type player struct {
	ID   int
	Name string
}

func (p *player) changeID(newID int) {
	p.ID = newID
}

func main() {
	fmt.Println("=====Ini merupakan contoh Slice and Struct di Golang=====")

	var players []player
	players = []player{player{1, "Tyo Marius"}, player{2, "Cristiano Ronaldo"}}

	players = append(players, player{3, "Lionel Messi"})

	for _, v := range players {
		fmt.Println(v)
	}

	for i, v := range players {
		if strings.EqualFold(v.Name, "Tyo Marius") {
			players[i].changeID(10)
			//v.changeID(10) --> setelah dicoba tidak bisa dikarenakan, v itu mengcopy value dari players[indeks] maka pasti addressnya juga berbeda.
			break
		}
	}

	fmt.Println("===Change value ID===")
	for _, v := range players {
		fmt.Println(v)
	}
}
