package main

import "fmt"

func main() {
	var nilai int
	
	fmt.Print("Masukkan nilai (0-100): ")
	fmt.Scan(&nilai)

	if nilai >= 90 {
		fmt.Println("Predikat: A")
	} else if nilai >= 70 {
	 	fmt.Println("Predikat: B")
	} else if nilai >= 60 {
	 	fmt.Println("Predikat: C")
	} else if nilai >= 50 {
	 	fmt.Println("Predikat: D")
	} else {
	 	fmt.Println("Predikat: E")
	}
}