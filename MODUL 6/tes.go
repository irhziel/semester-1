package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scan(&n)

	// Perulangan luar: mengatur jumlah baris
	for i := n; i >= 1; i-- {
		// Perulangan dalam: mencetak angka dari 1 hingga i
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// Jika n = 0, loop di atas tidak berjalan, tapi tetap cetak tanda "-"
	fmt.Println("-")
}
