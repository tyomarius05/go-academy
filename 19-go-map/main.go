package main

import (
	"fmt"
)

type player struct {
	ID   int
	Name string
}

func (p *player) getName() string {
	return p.Name
}

func main() {
	fmt.Println("=====Ini merupakan contoh Basic Map di Golang=====")

	fmt.Println()
	fmt.Println("Contoh pertama variabel Map")
	var mapPerson map[int]string
	mapPerson = make(map[int]string)

	// mapPerson[1] = "Tyo Marius"
	// mapPerson[2] = "Ancilla Intan"
	// mapPerson[5] = "Cristiano Ronaldo"
	// mapPerson[8] = "Lionel Messi"
	//bisa juga input data ke map dengan cara sbb
	mapPerson = map[int]string{
		1: "Tyo Marius",
		2: "Ancilla Intan",
		5: "Cristiano Ronaldo",
		8: "Lionel Messi",
	}

	mapPerson[9] = "Eden Hazard" //untuk menambahkan data baru

	for k, v := range mapPerson {
		fmt.Println(k, v)
	}

	//Contoh untuk pengecekan data apakah ada di map atau tidak
	v, isOK := mapPerson[4]
	if !isOK {
		fmt.Println("Indeks/key 4 tidak ditemukan")
	} else {
		fmt.Println(v)
	}

	fmt.Println()
	fmt.Println("Contoh kedua variabel Map atas Struct")

	// mapPlayer map[int]player
	mapPlayer := make(map[int]player)

	// mapPlayer[10] = player{ID: 1, Name: "Tyo Marius"}
	// mapPlayer[20] = player{ID: 2, Name: "Ancilla Intan"}
	// mapPlayer[30] = player{ID: 3, Name: "Cristiano Ronaldo"}
	// mapPlayer[40] = player{ID: 4, Name: "Lionel Messi"}

	mapPlayer = map[int]player{
		10: player{ID: 1, Name: "Tyo Marius"},
		20: player{ID: 2, Name: "Ancilla Intan"},
		30: player{ID: 3, Name: "Cristiano Ronaldo"},
		40: player{ID: 4, Name: "Lionel Messi"},
	}

	mapPlayer[50] = player{ID: 5, Name: "Romelu Lukaku"} //untuk menambahkan data baru

	for k, v := range mapPlayer {
		fmt.Println(k, v.getName())
	}
}
