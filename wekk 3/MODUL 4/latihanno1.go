package main

import "fmt"

func main() {
	var totalAwal int
	var diskon int

	fmt.Print("Total  Belanja Awal: ")
	fmt.Scan(&totalAwal)

	fmt.Print("Diskon (%): ")
	fmt.Scan(&diskon)

	totalAkhir := totalAwal - (totalAwal * diskon / 100)
	fmt.Printf("Total Belanja Akhir: %.d\n", totalAkhir)

}