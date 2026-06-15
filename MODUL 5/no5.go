package main

import "fmt"

func main() {
	var n int

	fmt.Print("Bilangan bulat positif: ")
	fmt.Scan(&n)

	jumlahDigit := 0
	for n > 0 {
		n = n / 10
		jumlahDigit++
	}

	fmt.Printf("Jumlah digit: %.d\n", jumlahDigit)
}