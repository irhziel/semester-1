package main

import "fmt"

func main() {
	var fahrenheit int
	var celcius float32

	fmt.Print("Masukkan suhu Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celcius = (float32(fahrenheit) - 32) * 5 / 9

	fmt.Printf("Suhu dalam Celcius = %.2f\n", celcius)

}