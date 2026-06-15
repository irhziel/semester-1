package main

import "fmt"

func main() {
    var n int             // deklarasi variabel n tipe integer
    fmt.Scan(&n)          // input dari user

    hasil := 0            

    for i := 1; i <= n; i++ {
        if n%i == 0 {     // cek apakah i faktor n

            prima := true // asumsikan bil. prima

            if i <= 1 {   // 0 dan 1 bukan prima
                prima = false
            }

            for j := 2; j*j <= i; j++ { // cek pembagi, kalau ada yang bagi habis berarti bukan bil. prima
                if i%j == 0 {  
					prima = false         
                }
            }

            if prima {           // kalau bil.prima, tambah hitungan
                hasil = hasil+1  
            }
        }
    }

    fmt.Println(hasil)   // cetak jumlah faktor prima
}