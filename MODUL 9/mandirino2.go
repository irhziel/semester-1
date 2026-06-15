package main

import "fmt"

func main() {
	var siswa, maks, bis int //deklarasi variabel tipe integer

	fmt.Scan(&siswa)
	fmt.Scan(&maks)

	bis = siswa / maks  //hitung jumlah bis penuh
	if siswa%maks != 0 {  //kalau ada sisa siswa, bus tambah 1 lagi
		bis++
	}
	fmt.Println(bis) //total bus
}