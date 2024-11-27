package main

import "fmt"

func main () {
	var jumlah, karung int
	fmt.Scan(&jumlah, &karung)
	karungs := (jumlah * karung + 999) / 1000
	fmt.Print(karungs , "karung")
}
