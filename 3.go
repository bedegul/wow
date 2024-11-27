package main

import "fmt"

func main() {
	var kode int
	fmt.Scan(&kode)
	
	var hasil bool
	hasil = (kode / 1000 == kode % 10)
	
	fmt.Printf("%t", hasil)
}1