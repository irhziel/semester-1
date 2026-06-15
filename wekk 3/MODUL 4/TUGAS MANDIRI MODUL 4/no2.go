package main

import "fmt"

func main() {
	// deklarasi variabel beserta tipe datanya
	var uts float64
	var uas float64
	var tugas float64

	// input nilai yang akan dimasukkan ke variabel
	fmt.Print("Masukkan nilai UTS: ")
	fmt.Scan(&uts)

	fmt.Print("Masukkan nilai UAS: ")
	fmt.Scan(&uas)

	fmt.Print("Masukkan nilai Tugas: ")
	fmt.Scan(&tugas)

	// menghitung nilai akhir dengan rumus
	nilaiAkhir := (uts * 0.3) + (uas * 0.4) + (tugas * 0.3)

	// menampilkan hasil perhitungan nilai akhir
	fmt.Printf("Nilai Akhir Mahasiswa: %.1f\n", nilaiAkhir)
}