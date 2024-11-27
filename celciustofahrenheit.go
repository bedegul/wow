package main

import "fmt"

func main () {
	var suhu int
	var satuan string
	fmt.Scan(&suhu, &satuan)
	if satuan == "Celcius" {
		hasilc := (suhu * 9 / 5) + 32
		fmt.Printf("Suhu dalam Fahrenheit adalah %d F", hasilc)
	} else if satuan == "Fahrenheit" {
		hasilf := (suhu - 32) * 5 / 9
		fmt.Printf("Suhu dalam Fahrenheit adalah %d C", hasilf)
		}
		
	
}