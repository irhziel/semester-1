package main

import "fmt"

func main() {
	var r int
	var hasil float32 = 3.14
	const phi = 3.14

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&r)

	hasil = phi * float32(r) * float32(r)
	fmt.Printf("Luas Lingkaran dengan jari-jari %d adalah %.2f\n")
	
}