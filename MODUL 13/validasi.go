package main

import "fmt"

func main() {
	var n, total, valid int
	var rata float64

	valid = 0
	total = 0
	for {
		fmt.Scan(&n)
		if n == -1 {
			break
		}
		if n > 0 {
			total = total + n
			valid++
		}
	}
	if valid == 0 {
		fmt.Println("0 0.00")
	} else {
		rata = float64(total) / float64(valid)
		fmt.Printf("%d %.2f\n", valid, rata)
	}
}