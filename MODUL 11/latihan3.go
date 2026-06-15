package main

import "fmt"

func main() {
	var warna string

	fmt.Scan(&warna)
	switch warna {
	case "Merah":
		fmt.Println("Warna keberanian")
	case "Biru":
		fmt.Println("Warna ketenangan")
	case "Hijau":
		fmt.Println("Warna kesuburan")
	case "Kuning":
		fmt.Println("Warna keceriaan")
	default:
		fmt.Println("Warna tidak dikenal")
	}
}