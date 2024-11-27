package main

import "fmt"

func main() {
    var tahun int
    fmt.Scan(&tahun)

    var hasil bool
    hasil = ( tahun % 4 == 0 ) && ( tahun % 5 == 0 ) && ( tahun % 100 != 0 )

    fmt.Printf("%t",hasil)
}
