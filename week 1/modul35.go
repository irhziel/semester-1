package main

import "fmt"

func main() {
	var x,y int
	fmt.Print("Masukkan nilai x dan y : ")
	fmt.Scan(&x,&y)
	fmt.Print(x % y)
}