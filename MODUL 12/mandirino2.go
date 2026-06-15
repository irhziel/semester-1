package main

import "fmt"

func main() {
	// deklarasi variabel tipe integer
	var angka int

	// input dari user
	fmt.Scan(&angka)

	// selama nilai dari variabel angka > 0, perulangan terus berjalan
	for angka > 0 {
		// ambil digit terahir angka dengan modulus 10
		digit := angka % 10
		// tampilkan digit terakhir
		fmt.Println(digit)
		// hapus digit terakhir dari angka dengan bagi 10
		angka = angka / 10
	}
}