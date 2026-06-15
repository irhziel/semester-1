package main

import "fmt"

func main() {
	// deklarasi variabel
	var nilai int

	// input dari user
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	// switch untuk cek range nilai
	switch {
	case nilai >= 85 && nilai <= 100:
		fmt.Printf("Grade Anda: A\n")
	case nilai >= 75 && nilai <= 84:
		fmt.Printf("Grade Anda: B\n")
	case nilai >= 65 && nilai <= 74:
		fmt.Printf("Grade Anda: C\n")
	case nilai >= 50 && nilai <= 64:
		fmt.Printf("Grade Anda: D\n")
	case nilai >= 0 && nilai < 50:
		fmt.Printf("Grade Anda: E\n")
	default:
		// jika diluar range 0-100, maka tidak valid
		fmt.Println("Tidak Valid")
	}
}