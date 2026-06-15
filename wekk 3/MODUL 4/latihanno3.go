package main

import(
	"fmt"
	"math"
)

func main() {
	var Ax, Ay float64
	var Bx, By float64
	var Cx, Cy float64

	fmt.Println("Titik A: ")
	fmt.Scan(&Ax, &Ay)

	fmt.Println("Titik B: ")
	fmt.Scan(&Bx, &By)

	fmt.Println("Titik C: ")
	fmt.Scan(&Cx, &Cy)

	AB := math.Sqrt(math.Pow(Bx-Ax, 2) + math.Pow(By-Ay, 2))
	BC := math.Sqrt(math.Pow(Cx-Bx, 2) + math.Pow(Cy-By, 2))
	CA := math.Sqrt(math.Pow(Ax-Cx, 2) + math.Pow(Ay-Cy, 2))

	titikTerpanjang := math.Max(AB, math.Max(BC, CA))
	fmt.Printf("Titik Terpanjang: %.2f\n", titikTerpanjang)
}