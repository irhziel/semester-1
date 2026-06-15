package main

import "fmt"

func main() {
	var jenis string
	var jam int

	fmt.Scan(&jenis, &jam)

	tarif := 0
	switch jenis {
	case "Mobil":
		switch {
		case jam <= 2:
			tarif = jam * 5000
		default:
			tarif = (2 * 5000) + (jam-2)*4000
		}

	case "Motor":
		switch {
		case jam <= 3:
			tarif = jam * 2000
		default:
			tarif = (3 * 2000) + (jam-3)*1500
		}

	case "Truk":
		switch {
		case jam <= 1:
			tarif = jam * 10000
		default:
			tarif = (1 * 10000) + (jam-1)*8000
		}

	default:
		fmt.Println("Tidak Valid")
		return
	}

	if jam > 8 {
		tarif = tarif - tarif/10
	}

	fmt.Printf("%d\n", tarif)
}