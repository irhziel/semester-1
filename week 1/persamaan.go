package main

import "fmt"

func main() {
	var x,y float32
	fmt.Print("Masukan x dan y: ")
	fmt.Scan(&x,&y)
	var hasil float32
	hasil=1/(3.0*x*x+10.0)+(10.0*y)+7.0

	fmt.Println(hasil)
}

