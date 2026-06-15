package main

import "fmt"

func main() {
	// deklarasi variabel tipe integer
	var x, y int

	// input dari user
	fmt.Scan(&x, &y)

	// cek apakah y bernilai 0
	if y == 0 {
		// jika y bernilai 0, tampilkan pesan tidak valid
		fmt.Println("Tidak Valid")
		// memberhentikan program
		return
	}
	// variabel menyimpan hasil pembagian
	hasil := 0
	// variabel sementara untuk modifikasi nilai x
	temp := x

	// selama variabel temp >= y, dilakukan perulangan
	for temp >= y {
		// mengurangi nilai nilai temp dengan y
		temp -= y
		// menghitung berapa kali pengurangan terjadi
		hasil++
	}

	// ouput hasil
	fmt.Println(hasil)
}