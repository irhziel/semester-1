package main

import "fmt"

func main() {
	var x float32

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	var fungsi float32 = 2 / (x - 5) - 5

	fmt.Printf("Hasil fungsi adaalah: %.1f", fungsi)
}
