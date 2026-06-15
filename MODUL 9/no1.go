package main

import "fmt"

func main() {
	var orang int
	var mobil int

	fmt.Scan(&orang)

	mobil = orang / 4
	if orang%4 != 0 {
		mobil++
	}

	fmt.Println(mobil)
}