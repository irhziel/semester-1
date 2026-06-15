package main

import "fmt"

func main() {
	// deklarasi variabel
	var bb float64
	var tb, bmi float64

	// input variabel
	fmt.Print("Masukkan berat (kg) : ")
	fmt.Scan(&bb)
	fmt.Print("Masukkan tinggi (m) : ")
	fmt.Scan(&tb)

	// hitung BMI
	bmi = bb / (tb * tb)

	// if-else dengan hasil output BMI dan kategori yang sesuai dengan yang user inputkan
	if bmi < 18.5 {
		fmt.Printf("BMI anda: %.2f (Kurus)\n", bmi)
	} else if bmi >= 18.5 && bmi < 25 {
		fmt.Printf("BMI anda: %.2f (Normal)\n", bmi)
	} else if bmi >= 25 && bmi < 30 {
		fmt.Printf("BMI anda: %.2f (Overweight)\n", bmi)
	} else if bmi >= 30 && bmi < 35 {
		fmt.Printf("BMI anda: %.2f (Obesitas I)\n", bmi)
	} else if bmi >= 35 && bmi < 40 {
		fmt.Printf("BMI anda: %.2f (Obesitas II)\n", bmi)
	} else if bmi >= 40 {
		fmt.Printf("BMI anda: %.2f (Obesitas III)\n", bmi)
	}

}