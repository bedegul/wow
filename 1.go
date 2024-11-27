package main

import "fmt"

func main() {
    var x, n int
    fmt.Print("Bilangan x dan N = ")
    fmt.Scan(&x, &n)

    var hasil bool
    hasil = (x % n == 0)

    fmt.Printf("%d kelipatan %d ? %t\n", x, n, hasil)
}
