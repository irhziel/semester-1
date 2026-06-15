package main

import "fmt"

func main() {
	var n int

	fmt.Print("Jumlah barang: ")
	fmt.Scan(&n)

	total := 0
	for i := 1; i <= n; i++ {
		var harga_satuan, jumlah_unit int

		fmt.Print("Harga satuan dan jumlah unit barang ke-%.d: ")
		fmt.Scan(&harga_satuan, &jumlah_unit)

		total += harga_satuan * jumlah_unit
	}

	fmt.Printf("Total biaya seluruh barang adalah %.d\n", total)
}