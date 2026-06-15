package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		nama, kelas string
		NIM int
	)

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan kelas: ")
	fmt.Scan(&kelas)

	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&NIM)

	fmt.Println("Perkenalkan saya adalah " + nama + ", salah satu mahasiswa Prodi S1-IF dari kelas " + kelas, "dengan NIM " + strconv.Itoa(NIM))
}