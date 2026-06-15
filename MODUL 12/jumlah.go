package main

import "fmt"

func main() {
	var n int

	total := 0

	fmt.Scan(&n)


	for n != 0 {
		total = total + n
		fmt.Scan(&n)
	}

	fmt.Printf("Total = %d\n", total)


}