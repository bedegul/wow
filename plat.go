package main

import "fmt"

func main() {
	var no, no1, no2, no3, no4, no5, jumlah, hasil int
	fmt.Scan(&no)
	
	if no > 99999 {
	fmt.Print("5 digit ya")
	}
	
	no1 = no%10
	no2 = (no/10)%10
	no3 = (no/100)%10
	no4 = (no/1000)%10
	no5 = (no/10000)%10
	
	jumlah = no1 + no2 + no3 + no4 + no5
	hasil = jumlah/10 + jumlah%10
	
	fmt.Println(hasil)
}