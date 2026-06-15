package main

import "fmt"

func main() {
	// deklarasi variabel
	var tanaman int
	//input
	fmt.Print("Masukkan jumlah tanaman: ")
	fmt.Scan(&tanaman)

	airEmber := 0.0
	isiEmber := 0
	penyiraman := 0

	for i := 0; i < tanaman; i++ {
		airTanaman := 0.0
		for airTanaman < 4 {
			// kalau ember kosong, isi
			if airEmber == 0 {
				airEmber = 3
				isiEmber++
			}
			// sekali nyiram, bocor 0.5
			if airEmber <= 0.5 {
				airEmber = 0
				continue
			}
			// air yang bisa dipakai
			airPakai := airEmber - 0.5
			// jangan lebih dari kebutuhan tanaman
			if airTanaman+airPakai > 4 {
				airPakai = 4 - airTanaman
			}
			airTanaman += airPakai
			airEmber -= airPakai
			penyiraman++
		}
	}
	//ouput
	fmt.Println("Total pengisian ember:", isiEmber, "kali")
	fmt.Println("Total penyiraman:", penyiraman, "kali")
}
