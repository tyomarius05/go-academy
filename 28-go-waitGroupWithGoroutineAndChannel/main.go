package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=====Ini merupakan contoh WaitGroup dengan Goroutine And Channel di Golang=====")
	fmt.Println()

	var wg sync.WaitGroup

	fmt.Println("=Contoh pertama=")
	wg.Add(1)

	go func() {
		fmt.Println("Proses selesai")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println()
	fmt.Println("=Contoh kedua=")
	var n int = 3
	wg.Add(n)

	// httpReq1 := <-process(1, 3, &wg)
	// httpReq2 := <-process(2, 1, &wg)
	// httpReq3 := <-process(3, 5, &wg)
	// fmt.Println(httpReq1)
	// fmt.Println(httpReq2)
	// fmt.Println(httpReq3)
	//dengan cara diatas ini maka semua akan ke print jika semuanya sudah selesai.

	httpReq1 := process(1, 3, &wg)
	httpReq2 := process(2, 1, &wg)
	httpReq3 := process(3, 5, &wg)
	for i := 1; i <= n; i++ {
		select {
		case p1 := <-httpReq1:
			fmt.Println(p1)
		case p2 := <-httpReq2:
			fmt.Println(p2)
		case p3 := <-httpReq3:
			fmt.Println(p3)
		}
	}
	wg.Wait()
}

func process(id int, delay time.Duration, wg *sync.WaitGroup) <-chan string {
	output := make(chan string)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * delay)

		output <- fmt.Sprintf("Proses ke-%d selesai", id)
	}()

	return output
}
