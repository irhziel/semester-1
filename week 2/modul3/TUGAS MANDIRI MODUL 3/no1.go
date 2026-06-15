package main

import "fmt"

func main() {
	// deklarasi variabel beserta tipe datanya
	var ms int
	var detik float64

	// input nilai yang akan dimasukkan ke variabel
	fmt.Print("Masukkan waktu (ms): ")
	fmt.Scan(&ms)

	// konversi ms ke detik
	detik = float64(ms) / 1000.0

	// menampilkan hasil konversi dari ms ke detik
	fmt.Printf("Waktu dalam detik: %.1f\n", detik)
}