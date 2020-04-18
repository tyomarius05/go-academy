package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Operator Logika dan IF di Golang=====")

	a := 200
	b := 200
	if a > b {
		fmt.Println("Nilai A =", a, "lebih besar dari pada nilai B =", b)
	} else if a < b {
		fmt.Println("Nilai A =", a, "lebih kecil dari pada nilai B =", b)
	} else {
		fmt.Println("Nilai A =", a, "sama dengan nilai B =", b)
	}

	/*
		Catatan:
		Pada saat membuat Syntax IF dan ELSE, maka antara } milik IF dengan syntax ELSE jangan beda baris.
		Dikarenakan Golang melakukan automatic insert semicolon (;) disetiap Enter atas syntax yang berkaitan dengan Ending.
	*/
}
