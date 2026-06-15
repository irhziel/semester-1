package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah bulan: ")
	fmt.Scan(&n) // membaca input jumlah bulan

	// inisialisasi dua nilai pertama deret Fibonacci
	a, b := 0, 1

	// cetak dua nilai pertama terlebih dahulu
	fmt.Print(a, " ", b, " ")

	// loop untuk menghitung sisa nilai Fibonacci
	for i := 3; i <= n; i++ {
		c := a + b          // jumlah dua angka sebelumnya
		fmt.Print(c, " ")   // cetak hasilnya
		a = b               // geser nilai
		b = c               // geser nilai
	}
}