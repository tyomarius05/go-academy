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

func (p player) getID() int {
	return p.ID
}

func (p player) getName() string {
	return p.PlayerName
}

func (p player) getPosition() string {
	return p.Position
}

func (p player) getOverallPoint() float32 {
	return p.OverallPoint
}

func (p *player) changePosition(newPos string) {
	p.Position = newPos
}

func main() {
	fmt.Println("=====Ini merupakan contoh Struct Method and Receiver di Golang=====")

	p := player{
		ID:           1,
		PlayerName:   "Tyo Marius",
		Position:     "CF",
		OverallPoint: 99.9,
	}

	printPlayer(p)

	p.changePosition("ST")

	printPlayer(p)

	fmt.Println("Contoh method reference")
	p2 := player{
		ID:           2,
		PlayerName:   "Yohanes",
		Position:     "AM",
		OverallPoint: 99.8,
	}

	printPlayer(p)
	printPlayer(p2)

	fmt.Println("==SWAP==")
	swapPlayer(&p, &p2)

	printPlayer(p)
	printPlayer(p2)
}

func printPlayer(p player) {
	fmt.Println("Player ID =", p.getID())
	fmt.Println("Player Name =", p.getName())
	fmt.Println("Player Position =", p.getPosition())
	fmt.Println("Player Overall Point =", p.getOverallPoint())
}

func swapPlayer(p1, p2 *player) {
	tempP := *p2
	*p2 = *p1
	*p1 = tempP
}
