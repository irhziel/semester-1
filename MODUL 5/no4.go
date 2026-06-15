package main

import "fmt"

func main() {
	var n int
	fmt.Print("Bilangan bulat: ")
	fmt.Scan(&n)

	var jumlah float64

	for i := 0; i < n; i++ {
		var nilai float64
		fmt.Scan(&nilai)
		jumlah += nilai
	}

	rataRata := jumlah / float64(n)
	fmt.Printf("%.2f\n", rataRata)
}