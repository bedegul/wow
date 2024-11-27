package main

import "fmt"

func main () {
	var x int
	
	fmt.Scan(&x)
	
	switch {
	case x >= 1 && x <= 9 :
		fmt.Print("Satuan")
	case x >= 10 && x <= 99 :
		fmt.Print("Puluhan")
	case x >= 100 && x <= 999 :
		fmt.Print("Ratusan")
	case x >= 1000 && x <= 9999 :
		fmt.Print("Ribuan")
	case x >= 10000 && x <= 99999 :
		fmt.Print("Puluhan ribu")
	case x >= 100000 && x <= 999999 :
		fmt.Print("Ratusan ribu")
	default :
		fmt.Print("Kiminosei kiminosei")
	}
}