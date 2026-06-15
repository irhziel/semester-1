package main

import (
    "fmt"
    "strconv"
)

func main() {
    var teks string = "2024"
    tahun, err := strconv.Atoi(teks)
    if err == nil {
        fmt.Println(tahun + 2) // output: 2024
    }
}