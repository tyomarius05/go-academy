package main

import (
	"encoding/json"
	"fmt"
)

type player struct {
	ID   int    `json:"id"` //`json:` digunakan untuk custom key saja di JSON
	Name string `json:"name"`
}

func main() {
	fmt.Println("=====Ini merupakan contoh JSON Package di Golang=====")
	fmt.Println()

	fmt.Println("==Contoh Encode dengan Marshal==")
	myPlayer := player{ID: 1, Name: "Tyo Marius"}

	jsonPlayer, err := json.Marshal(myPlayer) //Marshal ini seperti encode

	if err != nil {
		fmt.Println("Terjadi kesalahan:", err.Error())
	} else {
		fmt.Println(string(jsonPlayer))
	}

	fmt.Println()
	fmt.Println("==Contoh Decode dengan UnMarshal==")

	data := []byte(`{"id":1, "name":"Lionel Messi"}`)

	var myPlayer2 player

	err = json.Unmarshal(data, &myPlayer2) //Unmarshal memiliki 2 argument yaitu byte dan tipe data struct. Fungsi Unmarshal untuk decode JSON

	if err != nil {
		fmt.Println("Terjadi kesalahan:", err.Error())
	} else {
		fmt.Println(myPlayer2.ID)
		fmt.Println(myPlayer2.Name)
	}
}
