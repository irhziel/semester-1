package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	angka := 2

	for angka <= n {
		fmt.Print(" ",angka)
		angka += 2
		
	}


}