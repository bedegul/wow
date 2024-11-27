package main

import "fmt"

func main () {
	var digit int
	fmt.Scan(&digit)
	no1 := digit/100
	no2 := (digit/10)%10
	no3 := digit%10
	fmt.Printf("%d%d%d%d%d%d",no1,no1,no2,no2,no3,no3)
}