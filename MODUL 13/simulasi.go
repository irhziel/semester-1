package main

import "fmt"

func main() {
	var n, digit int

	digit = 0
	fmt.Scan(&n)
	if n == 0 {
		fmt.Println(1)
		return
	}
	for n != 0 {
		n = n / 10
		digit++
	}
	fmt.Println(digit)

}