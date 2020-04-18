package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Switch-Case di Golang=====")

	var x int
	x = 2

	//Contoh switch tipe pertama
	switch x {
	case 1:
		fmt.Println("x = satu")
		break
	case 2:
		fmt.Println("x = dua")
	case 3:
		fmt.Println("x = tiga")
		break
	case 4:
		fmt.Println("x = empat")
		break
	case 5:
		fmt.Println("x = lima")
		break
	default:
		fmt.Println("x = tidak diketahui")
		break
	}

	//Switch Case tipe kedua
	switch {
	case x == 1:
		fmt.Println("x = satu")
		break
	case x == 2:
		fmt.Println("x = dua")
	case x == 3:
		fmt.Println("x = tiga")
		break
	case x == 4:
		fmt.Println("x = empat")
		break
	case x == 5:
		fmt.Println("x = lima")
		break
	default:
		fmt.Println("x = tidak diketahui")
		break
	}

	var a interface{}
	a = 1

	switch a.(type) {
	case int:
		fmt.Println("variable a bertipe integer")
		break
	case string:
		fmt.Println("variable a bertipe string")
	case float64:
		fmt.Println("variable a bertipe float64")
		break
	default:
		fmt.Println("x = tidak diketahui")
		break
	}
}
