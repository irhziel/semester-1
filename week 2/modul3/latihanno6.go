package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari float64

	fmt.Scan(&jariJari)

	volume := (4.0 / 3.0) * math.Pi *  math.Pow(jariJari, 3)
	luas := 4 * math.Pow(jariJari, 2)

	fmt.Print(volume, luas)
}