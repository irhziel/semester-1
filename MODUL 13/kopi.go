package main
import "fmt"

func main() {
    var pass string

    for {
        fmt.Print("Masukkan password: ")
        fmt.Scan(&pass)

        if pass == "kopi123" {
            break
        }
        fmt.Println("Salah! Coba lagi.")
    }

    fmt.Println("Login berhasil!")
}