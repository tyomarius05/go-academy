package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Slice di Golang=====")
	fmt.Println()

	fmt.Println("Contoh pertama, Slice sebagai array tanpa mendefinisikan jumlah value yang dapat ditampung")

	mySlice := []int{1, 2, 3}             //kita bisa menginputkan berapa banyak valuenya, tanpa mendeklarasi jumlah value yg dapat ditampung
	mySlice = append(mySlice, 4, 5, 6, 7) //untuk menambahkan valuenya

	for _, v := range mySlice {
		fmt.Println(v)
	}

	fmt.Println("Contoh kedua, slice sebagai pointer dari var array")
	//slice bisa digunakan untuk mempointerkan diri ke suatu var array
	//jika nilai slice diubah maka akan mempengaruhi nilai dari var array

	myArr := [5]int{10, 20, 30, 40, 50}
	fmt.Println(myArr)   //mencetak nilai di var array
	slice1 := myArr[1:4] //mempointerkan ke var array untuk indeks ke 1 sampai 4
	fmt.Println(slice1)  //jika dicetak akan menampilkan angka 20, 30, dan 40

	slice1[1] = 300 //isian slice tadi indeks 0 = 20, indeks 1 = 30, dan indeks 2 = 40. Disini kita akan mengubah value slice indeks ke 1 yaitu 30 menjadi 300
	//akibatknya akan mengubah nilai pada var indeks ke 2 yaitu 30 menjadi 300 juga. Hal ini karena slice itu mempointerkan diri ke var array pada indeks tertentu

	fmt.Println(slice1)
	fmt.Println(myArr)
	fmt.Println()

	fmt.Println("Contoh ketiga, slice")
	slice2 := make([]string, 3) //cara lainnya membuat slice array, dengan banyak value yang ditampung sebanyak 3 indeks
	slice2[0] = "Tyo"
	slice2[1] = "Marius"
	slice2[2] = "Golang"

	slice2 = append(slice2, "Bisa") //menambahkan value ke slice array
	fmt.Println(slice2)

	slice2[1] = "Wahyu" //jika suatu indeks di slice2 diubah maka slice1 juga akan berubah, karena append itu tetep merefer/mempointerkan
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := append(slice2, "NewSliceArr") //append bisa digunakan untuk mengcopy nilai yg ada di slice2 ke slice3 dengan menambahkan 1 value lagi yaitu "NewSliceArr"
	//dan jika ada perubahan pada slice 3 maka akan mengubah nilai di slice yang dicopy/slice2
	fmt.Println(slice3)
	slice3[0] = "Yohanes"
	fmt.Println(slice2)
	fmt.Println(slice3)

	slice4 := make([]string, 3)
	copy(slice4, slice3) //cara mengcopy value dari suatu slice ke slice lain dengan tidak saling mempointerkan diri
	//dan data yang dapat ditampung dan dicopy oleh slice4 adalah sebanyak 3, karena disaat deklarasi jumlah valuenya 3

	fmt.Println(slice4)
	slice4[0] = "Tyokkk"
	fmt.Println(slice4)
	fmt.Println(slice3)

	var slice5 []string //dengan cara seperti ini apakah bisa mengcopy slice?
	copy(slice5, slice3)

	fmt.Println(slice5) //setelah ditest ternyata tidak bisa, karena syntax copy tidak mengcopy jumlah indeksnya, tapi hanya tercopy valuenya saja.
	//jika slice5 tidak dideskripsikan jumlah indeksnya maka dianggap bahwa slice5 hanya bisa menampung nil indeks.
}
