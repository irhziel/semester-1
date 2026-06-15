package main

import (
	"fmt"
	"math"
)

func main() {
	var nilaiBMI float64
	var tinggiBadan float64

	fmt.Print("Nilai BMI: ")
	fmt.Scan(&nilaiBMI)

	fmt.Print("Tinggi Badan (m): ")
	fmt.Scan(&tinggiBadan)

	beratBadan := nilaiBMI * math.Pow(float64(tinggiBadan), 2)

	fmt.Printf("Berat Badan: %.f kg", beratBadan)
}