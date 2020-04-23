package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Environment Variable di Golang=====")
	fmt.Println()

	t := time.Now()

	fmt.Println(t)
	fmt.Println(t.Day())
	fmt.Println(t.Month())
	fmt.Println(t.Year())

	fmt.Println("===============")

	timeString := "January 22, 2020"
	format := "January 2, 2006"

	resTime, err := time.Parse(format, timeString)

	if err != nil {
		fmt.Println("Terjadi kesalahan:", err.Error())
	}

	fmt.Println(resTime)
	fmt.Println(resTime.Year())

	fmt.Println("===============")

	t1 := time.Date(2020, 2, 1, 12, 0, 0, 0, time.UTC)

	fmt.Println(t1)

	t2 := time.Date(2020, 2, 1, 12, 0, 0, 0, time.UTC)

	fmt.Println(t1.Equal(t2))

	fmt.Println("===============")

	layout := "2006-01-02"

	timeNow := time.Now()

	resString := timeNow.Format(layout)

	fmt.Println(resString)
}
