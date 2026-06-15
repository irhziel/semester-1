package main

import "fmt"

func main() {
	var nilai int
	var grade string
	var pred string

	fmt.Scan(&nilai,&grade,&pred)

	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"
		pred = "Excellent"
		if nilai%2 == 1 {
			fmt.Println("Outstanding")
		}
	case nilai >= 80 && nilai <= 89:
		grade = "B"
		pred = "Very Good"
	case nilai >= 70 && nilai <= 79:
		grade = "C"
		pred = "Good"
	case nilai >= 60 && nilai <= 69:
		grade = "D"
		pred = "Satisfactory"
		if nilai%2 == 0 {
			fmt.Println("")
		}
	case nilai >= 50 && nilai <= 59:
		grade = "E"
		pred = "Very Bad"
	}
}