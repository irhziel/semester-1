package main

import "fmt"

func main() {
	var suhu int
	fmt.Print("Masukkan suhu dalam Fahrenheit : ")
	fmt.Scan(&suhu)
	celcius := (suhu - 32) * 5 / 9
	fmt.Printf("Suhu dalam celcius: %d", celcius)
}