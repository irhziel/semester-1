package main

import "fmt"

func main() {
	var n float64

	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("0.00000")
		return
	}
	x := n 
	xbaru := 0.0
	for {
		xbaru = (x + n / x) / 2
		selisih := xbaru - x
		if selisih < 0 {
			selisih = -selisih
		}
		if selisih < 0.00001 {
			break
		}
		x = xbaru

	}
	fmt.Printf("%.5f\n", xbaru)
}