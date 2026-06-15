
package main

import (
    "fmt"
)

func main() {
    pwBenar := "kopi123"
    percobaan := 0
    maxCoba := 3
    var pwInput string

    for percobaan < maxCoba {
        fmt.Print("Input password: ")
        fmt.Scanln(&pwInput)

        if pwInput == pwBenar {
            fmt.Println("Login sukses!")
            return
        } else {
            percobaan++
            if percobaan < maxCoba {
                fmt.Printf("Salah, coba lagi (sisa %dx)\n", maxCoba-percobaan)
            }
        }
    }

    fmt.Println("Akses ditolak!")
}
