package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=====Ini merupakan contoh WorkerPool Goroutine And Channel (Producer and Consumer) di Golang=====")
	fmt.Println()

	start := time.Now()

	jobs := make(chan int, 100)
	results := make(chan int, 100)
	fmt.Println("=====================")

	go consumer(1, jobs, results) //Consumer disini untuk mengisikan nilai ke channel Results

	var n int = 10
	go producer(jobs, n) //Producer disini akan mengisikan jumlah data jobs yang dapat diiterasikan oleh Consumer

	for i := 1; i <= n; i++ {
		res := <-results //Perulangan akan jalan jika ada tanda dari result atau jika result terisi (Alur cerita akan dijelaskan dibawah)
		fmt.Println("Hasil ke", res)
	}

	elapsed := time.Since(start)
	fmt.Println("=====================")
	fmt.Println("Waktu:", elapsed)
}

func fakeHTTPRequest(x int) int {
	return x * 10
}

func producer(job chan<- int, size int) {
	for i := 1; i <= size; i++ {
		job <- i
	}
	fmt.Println("Data sudah dimasukkan ke Job sebanyak:", len(job))
	close(job)
}

func consumer(id int, jobs <-chan int, results chan<- int) {
	fmt.Println("Jumlah jobs yang akan diproses", len(jobs))
	for job := range jobs {
		fmt.Println("Consumer ke", job, "Mulai, Job yg tersisa", len(jobs))

		time.Sleep(time.Second * 2)

		results <- fakeHTTPRequest(job)
	}
}

/*
	Catatan: bahwa Main itu merupakan suatu goroutine juga, jadi ketika di dalam Main itu ada menjalankan goroutine lainnya maka Main tidak
	akan menunggu sampai goroutine lainnya selesai, kecuali memang ada penandanya yang mengharuskan untuk menunggu. Penandanya bisa dilihat
	pada tutorial sebelumnya, atau jika disini adalah Results.

	Jadi alurnya adalah, Main akan menjalankan syntax go consumer terlebih dahulu kemudian tanpa menunggu go consumer maka Main akan
	menjalankan go producer. Dan tanpa menunggu lagi Main menjalankan perulangan untuk menampilkan nilai Results.
	Tetapi karena nilai Results belum ada, maka Main akhirnya menunggu.

	Jika dicek Go Customer adalah syntax di awal, tetapi jika dijalankan malah Go Producer duluanlah yang jalan (belum ketemu alasannya).
	Kemudian ketika Go Producer sudah mengisikan jumlah Job, maka Go Customer mulai bisa mengiterasi Job-job, sebanyak 10 kali (sesuai nilai n di Main).
	Lalu ketika sudah selesai sleep, maka results diisi nilainya.
	Karena nilai Results sudah keiisi 1 nilai, maka di Main bagian perulangan yang sebelumnya terhenti menjadi bisa jalan lagi dan
	mengeluarkan nilai Results, karena Results sudah ada nilainya.
	Kemudian perulangan berjalan lagi dan terhenti lagi menunggu nilai Results.
	Setelah itu Go Customer setelah sleep 2 detik maka mengisikan lagi nilai Results, setelah itu perulangan di Main bisa jalan lagi.
	Dan begitu seterusnya hingga teriterasi sebanyak 10 kali.
*/
