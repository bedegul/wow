package main

import "fmt"

func main () {
	var n int
	fmt.Scan(&n)
	
	if n >= 1 && n <= 9 {
		fmt.Println("Satuan")
	} else if n >= 10 && n <= 99 {
		fmt.Println("Puluhan")
	} else if n >= 100 && n <= 999 {
		fmt.Println("Ratusan")
	} else if n >= 1000 && n <= 9999 {
		fmt.Println("Ribuan")
	} else if n >= 10000 && n <= 99999 {
		fmt.Println("Puluhan ribu")
	} else if n >= 100000 && n <= 999999 {
		fmt.Println("Ratusan ribu")
	}
}
