package main

import "fmt"

func main () {
	var KI, SI bool
	fmt.Scan(&KI, &SI)
	
	if KI || SI {
	fmt.Println("Diizikan masuk")
	} else {
	fmt.Println("Tidak diizinkan masuk")
	}
}
