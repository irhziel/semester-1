package main

import "fmt"

func main() {
	var harga float64
	var mataUang, pelanggan string

	fmt.Scan(&harga)
	fmt.Scan(&mataUang)
	fmt.Scan(&pelanggan)

	pajak := 0.0
	diskon := 0.0

	switch mataUang {
	case "USD":
		if harga > 1000 {
			diskon = 0.05
		}

		switch pelanggan {
		case "Regular":
			pajak = 0.10
		case "Member":
			pajak = 0.05
		case "VIP":
			pajak = 0.0
		}
	case "EUR":
		if harga > 900 {
			diskon = 0.07
		}

		switch pelanggan {
		case "Regular":
			pajak = 0.10
		case "Member":
			pajak = 0.05
		case "VIP":
			pajak = 0.0
		}
	case "JPY":
		if harga > 120000 {
			diskon = 0.03
		}
		switch pelanggan {
		case "Regular":
			pajak = 0.10
		case "Member":
			pajak = 0.05
		case "VIP":
			pajak = 0.0
		}
	case "IDR":
		if harga > 10000000 {
			diskon = 0.02
		}

		switch pelanggan {
		case "Regular":
			pajak = 0.10
		case "Member":
			pajak = 0.05
		case "VIP":
			pajak = 0.0
		}
	default:
		fmt.Println("Tidak Valid")
		return
	}

	hasil := harga + (harga * pajak)
	hasil = hasil - (hasil * diskon)

	fmt.Printf("Total: %.2f %s\n", hasil, mataUang)
}