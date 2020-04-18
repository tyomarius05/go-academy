package main

import (
	"fmt"
)

func main() {
	var a int //static variable
	a = 1     //set value static variable
	fmt.Println(a)

	b := 2 //dynamic variable and set value, jadi tipe data pada dynamic variable itu akan tergantung dari value-nya di awal.
	//Jika awalnya valuenya bertipe integer maka si variable akan bertipe integer selamanya
	fmt.Println(b)

	var c string = "Ini merupakan contoh variable string"
	fmt.Println(c)

	d := "Ini contoh Dynamic Variable yang bertipe data String"
	fmt.Println(d)
}
