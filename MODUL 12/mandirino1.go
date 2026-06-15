package main

import "fmt"

func main() {
	// deklarasi variabel tipe string
	var usname, pw string
	// membuat variabel penghitung percobaan login
	percobaan := 0

	// infinite loop
	for { 
		// input dari user
		fmt.Print("Masukkan Username: ")
		fmt.Scan(&usname)
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&pw)

		// cek apakah username dan password benar
		if usname == "Admin" && pw == "Admin" {
			// jika benar, maka akan keluar dari loop
			break
		}
		// jika salah, maka jumlah percobaan gagal +1
		percobaan++
		// pesan jika gagal 
		fmt.Println("Login gagal, coba lagi")
	}
	// output jumlah percobaan gagal login
	fmt.Printf("%d perrcobaan gagal login\n", percobaan)
}