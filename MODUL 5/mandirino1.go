package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n) // membaca input angka dari pengguna

	//ubah angka ke bentuk string agar mudah dibalik
	str := strconv.Itoa(n)

	// variabel untuk menyimpan hasil pembalikan string
	reversed := ""

	// melakukan pembalikan string
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i]) // ambil karakter dari belakang satu per satu 
	}

	// hasilnya berupa nilai boolean, true jika palindrome dan false jika tidak
	fmt.Println(str == reversed)
}