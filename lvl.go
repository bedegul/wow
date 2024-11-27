package main

import "fmt"

func main () {
	var lvl int
	
	fmt.Scan(&lvl)
	
	switch {
	case lvl >= 1 && lvl <= 10 :
		fmt.Print("Pemula")
	case lvl >= 11 && lvl <= 20 :
		fmt.Print("Menengah")
	case lvl >= 21 && lvl <= 30 :
		fmt.Print("Ahli")
	case lvl > 30 :
		fmt.Print("Master")
	default :
		fmt.Print("Level tidak valid")
	}
}