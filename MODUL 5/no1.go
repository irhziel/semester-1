package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	jumlah := 0
	for i := 1; i <= n; i++ {
		jumlah += ( i % 2) * i
	}

	fmt.Printf("Hasil penjumlahan bilangan ganjil dari 1 hingga %.d adalah: %.d\n", n, jumlah)
}