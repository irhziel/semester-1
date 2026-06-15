package main

import "fmt"

func main() {
	var nilai1,nilai2,nilai3 int
	var total int
	var ratarata float64

	fmt.Print("Masukkan tiga niai: ")
	fmt.Scan(&nilai1,&nilai2,&nilai3)

	total = nilai1 + nilai2 + nilai3
	ratarata = float64(total) / 3

	fmt.Println("Total = ", total)
	fmt.Println("Rata-rata = ", ratarata)
}