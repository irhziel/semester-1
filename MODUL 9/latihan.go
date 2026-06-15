package main

import "fmt"

func main() {
	var bilangan int

	fmt.Scan(&bilangan)

	if bilangan > 0 {
		fmt.Println("Bilangan positif")
	}
	if bilangan == 0 {
		fmt.Println("Bilangan Nol")
	}
	if bilangan < 0 {
		fmt.Println("Bilangan Negatif")
	}
}