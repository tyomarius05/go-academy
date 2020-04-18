package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Pointer di Golang=====")

	x := "Tyo" //variable string biasa
	y := &x    //variable y mencatat address dari variable x
	fmt.Println(x)
	fmt.Println(y, &y) //variable y berisi address variable x, dan &y merupakan address dari variable y
	var z *string = &x //membuat variable pointer bernama z yang berisi alamat dari variable x
	*z = "Tyo Marius"  //mengubah isi pointer dari variable z untuk mengubah nilai dari x
	fmt.Println(x)

	nilA := 123
	nilB := 321
	fmt.Println(nilA, nilB)
	swap(&nilA, &nilB) //contoh penggunaan pointer untuk mengubah nilai dari variable yang dikirimkan ke function
	fmt.Println(nilA, nilB)
}

func swap(a, b *int) { //ini contoh pointer sebagai reference
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}
