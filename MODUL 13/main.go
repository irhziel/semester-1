package main
import "fmt"

func main() {
    var x int

    for {
        fmt.Scan(&x)
        fmt.Println("Anda memasukkan:", x)

        // kondisi berhenti: jika x < 0
        if x < 0 {
            break
        }
    }

    fmt.Println("Program selesai")
}