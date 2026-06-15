package main

import "fmt"

func main() {
	// deklarasikan variabel
	var n int

	fmt.Scan(&n)

	// gunakan for loop untuk menghitung mundur dari n ke 1
	for i := n; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println("")
}