package main

import "fmt"

func main() {
	// deklarasi variabel beserta tipe datanya
	var namaKaryawan string
	var gajiPokok float64
	var pajak float64

	// input nilai yang akan dimasukkan ke variabel
	fmt.Print("Masukkan nama karyawan: ")
	fmt.Scan(&namaKaryawan)

	fmt.Print("Masukkan gaji pokok: ")
	fmt.Scan(&gajiPokok)

	fmt.Print("Masukkan persentase pajak (%): ")
	fmt.Scan(&pajak)
	
	// menghitung pajak dan gaji bersih dengan rumus
	Pajak := (pajak / 100) * gajiPokok
	gajiBersih := gajiPokok - (pajak / 100) * gajiPokok

	// menampilkan hasil perhitungan pajak dan gaji bersih
	fmt.Printf("Nama Karyawan: %.s\n", namaKaryawan)
	fmt.Printf("Gaji Pokok: %.f\n", gajiPokok)
	fmt.Printf("Pajak: %.f\n", Pajak)
	fmt.Printf("Gaji Bersih: %.f\n", gajiBersih)
}