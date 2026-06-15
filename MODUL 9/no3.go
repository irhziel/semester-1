package main

import "fmt"

func main() {
	var x, y  int
	var xkely bool
	var ykelx  bool

	fmt.Scan(&x, &y)

	if x%y == 0 {
		xkely = true
	}
	if y%x == 0 {
		ykelx = true
	}
	fmt.Println(xkely)
	fmt.Println(ykelx)


}