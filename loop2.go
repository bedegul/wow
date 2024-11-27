package main

import "fmt"

func main () {
	var sgtg int
	var alas, tinggi float64
	var hasil float64
	fmt.Scan(&sgtg)
	
	for i := 0; i < sgtg; i++ {
		fmt.Scan(&alas, &tinggi)
		hasil = 0.5 * alas * tinggi
		fmt.Printf("%.2f", hasil)
		}
}