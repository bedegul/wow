package main

import "fmt"

func main () { 
	var uang float64
	fmt.Scan(&uang)
	
	for i := 1; i <= 10; i++ {
	uang = uang * 1.05
	fmt.Printf("Tahun ke- %d : %.2f\n", i , uang)
	}
}