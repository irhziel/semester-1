package main

import (
	"fmt"
)


func main() {
	food := 1

	kenyang := 20

	for food <= kenyang {
		fmt.Println("ini makanan saya ke-", food)
		food = food + 1 // <- stopper

	}
	fmt.Println("wareg lur!!")
}
