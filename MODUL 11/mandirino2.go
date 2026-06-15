package main

import "fmt"

func main() {
	// deklrasi variabel
	var kodeKendaraan, durasi int
	var tarif int
	var jenis string

	// input dari user
	fmt.Print("Masukkan kode kendaraan (1-4): ")
	fmt.Scan(&kodeKendaraan)
	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scan(&durasi)

	// jenis kendaraan dan tarif per jam berdasarkan kode
	switch kodeKendaraan {
	case 1:
		jenis = "Motor"
		tarif = 2000
	case 2:
		jenis = "Mobil"
		tarif = 5000
	case 3:
		jenis = "Truk"
		tarif = 8000
	case 4:
		jenis = "Bus"
		tarif = 10000
	}

	// total biaya parkir
	total := 0
	if durasi <= 5 {
		total = durasi * tarif
	} else {
		// tarif normal 5 jam pertama
		total = 5 * tarif
		// setelah 5 jam, tarif 1.5x
		tarifExtra := durasi - 5
		total += tarifExtra * int(float64(tarif)*1.5)
	}

	if durasi <= 0 {
		fmt.Println("Input tidak valid")
	}

	// output
	fmt.Println("Jenis Kendaraan :", jenis)
	fmt.Println("Durasi Parkir :", durasi, "jam")
	fmt.Println("Tarif per Jam :", "Rp", tarif)
	fmt.Println("Total Biaya Parkir :", "Rp", total)
}
