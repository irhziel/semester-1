package main

import "fmt"

func main() {
	// deklarasi variabel beserta tipe datanya
	var A, B, C float64
	var volumeTabung float64
	var volumeKerucut float64
	var volumeTotal float64	
	const phi = 3.14

	// input nilai yang akan dimasukkan ke variabel
	fmt.Print("Masukkan jari-jari alas: ")
	fmt.Scan(&A)
	
	fmt.Print("Masukkan tinggi tabung: ")
	fmt.Scan(&B)

	fmt.Print("Masukkan tinggi kerucut: ")
	fmt.Scan(&C)

	// menghitung volume tabung, volume kerucut, dan volume total
	volumeTabung = phi * A * A * B
	volumeKerucut = (1.0/3.0) * phi * A * A * C
	volumeTotal = volumeTabung + volumeKerucut

	// menampilkan hasil perhitungan volume tabung, volume kerucut, dan volume total
	fmt.Printf("Volume Tabung: %.2f\n", volumeTabung)
	fmt.Printf("Volume Kerucut: %.2f\n", volumeKerucut)
	fmt.Printf("Volume Total: %.2f\n", volumeTotal)
}