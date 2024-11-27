package main

import "fmt"

func main () {
	var uang int
	fmt.Scan(&uang)
	
	if uang > 1000 {
		uang = uang * 80/100
		fmt.Print(uang)
	} else if uang > 500 && uang < 1000 {
		uang = uang * 85/100
		fmt.Print(uang)
	} else {
		uang = uang * 95/100
		fmt.Print(uang)
	}
}