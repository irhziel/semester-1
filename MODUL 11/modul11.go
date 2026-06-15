package main

import "fmt"

func main() {
    var jam24 int
    var jam12 int
    var label string

    fmt.Print("Masukkan jam (0-23): ")
    fmt.Scan(&jam24)

    switch {
    case jam24 == 0:
        jam12 = 12
        label = "AM"
    case jam24 < 12:
        jam12 = jam24
        label = "AM"
    case jam24 == 12:
        jam12 = 12
        label = "PM"
    case jam24 < 12:
        jam12 = jam24 - 12
        label = "PM"
	default:
		break
    }

	if jam12 > 24 {
		fmt.Println("Jam Tidak Valid")
	} else {
		fmt.Println(jam12, label)
	}

    fmt.Println(jam12, label)
}