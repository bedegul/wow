package main

import "fmt"

func main() {
    var x, y, total int
    fmt.Scan(&x, &y)

    for i := x; i <= y; i++ {
        fmt.Print(i, " ")
        total += i
    }

    fmt.Println()
    fmt.Println("Total:", total)
}