package main

import "fmt"

func main() {
	var a,b float32
	fmt.Print("Masukkan nilai a dan b: ")
	fmt.Scan(&a,&b)
	fmt.Print(a / b)
}