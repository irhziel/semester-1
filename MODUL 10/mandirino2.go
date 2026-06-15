package main

import "fmt"

func main() {
	// deklarasi variabel
	var jenis string
	var jarak, tarif, diskon, tarifNrml, total float64

	// input variabel
	fmt.Print("masukkan Jenis Kendaraan (Mobil/Bus/Truk): ")
	fmt.Scan(&jenis)
	fmt.Print("Masukkan Jarak Tempuh (km): ")
	fmt.Scan(&jarak)

	// if-else tarif per km berdasarkan jenis kendaraan
	if jenis == "Mobil" {
		tarif = 1000
	} else if jenis == "Bus" {
		tarif = 1500
	} else if jenis == "Truk" {
		tarif = 2000
	} else {
		fmt.Println("")
	}

	// jika tarif tidak 0 maka dilakukan perhitungan tarif normal
	if tarif != 0 {
		tarifNrml = tarif * jarak
	}

	// if-else menentukan diskon sesuai jarak
	if jarak >= 100 {
		diskon = 25
	} else if jarak >= 50 && jarak < 100 {
		diskon = 15
	} else if jarak >= 20 && jarak < 50 {
		diskon = 5
	} else {
		diskon = 0
	}

	// hitung total bayar
	total = tarifNrml * (1 - diskon / 100)
	// output
	fmt.Println("Jenis Kendaraan: ", jenis)
	fmt.Printf("jarak Tempuh: %.f km\n", jarak)
	fmt.Printf("Tarif Normal: Rp %.f\n", tarifNrml)
	fmt.Printf("Diskon: %.f %%\n", diskon)
	fmt.Printf("Total Bayar: Rp %.f\n", total)
}