package main

import "fmt"

func main () {
	var adik, kaka int
	var hasil bool
	fmt.Scan(&adik, &kaka)
	hasil = adik == kaka || adik - kaka == 1 || adik - kaka == 5 || kaka - adik == 1 || kaka - adik == 5
	fmt.Print(hasil)
}