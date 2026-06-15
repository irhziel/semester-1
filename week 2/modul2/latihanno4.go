package main

import "fmt"

func main() {
	var (
		fahrenheit, celcius int
	)

	fmt.Print("Masukkan suhu Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celcius = (fahrenheit - 32) * 5 / 9
	fmt.Println("Suhu dalam celcius adalah", celcius)
}