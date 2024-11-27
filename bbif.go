package main

import "fmt"

func main () {
	var bb, tb float64
	fmt.Scan(&tb, &bb)
	
	tb = tb/100
	hasil := bb / (tb * tb)
	
	if hasil >= 18.5 && hasil <= 22.9 {
		fmt.Print("Berat badan normal")
	} else if hasil < 18.5 {
		fmt.Print("Berat badan kurang")
	} else if hasil > 22.9 {
	 fmt.Print("Kelebihan berat badan")
	 }
}