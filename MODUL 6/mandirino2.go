package main

import "fmt"

func main() {
	// deklarasikan variabel
	var n int

	fmt.Scan(&n)

	// mengatur berapa baris segitiga akan dicetak (berkurang satu tiap baris)
	for i := n; i >= 1; i-- {
		// mencetak angka dari 1 sampai i pada baris tersebut
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	fmt.Println("")
}