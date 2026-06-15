package main

import "fmt"

func main() {
	// deklarasi variabel
	var angka int
	// input
	fmt.Println("Tebak Angka Riri vs Roro")
	fmt.Scan(&angka)

	// loop tebakan dari 1 sampai 10
	for tebak := 1; tebak <= 10; tebak++ {
		// menampilkan tebakan Roro
		fmt.Printf("Roro Menjawab: %d\n", tebak)
		// cek apakah tebakan benar
		if tebak == angka {
			// jika benar
			fmt.Println("Riri Menjawab: Kamu Benar!!!")
			break // hentikan loop
		} else {
			// jika salah
			fmt.Println("Riri Menjawab: kamu Salah!!!")
		}
	}
}
