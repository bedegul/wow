package main

import "fmt"

func main () {
	var n, m, num int
	var benar = true
	
	fmt.Scan(&n, &m)
	for i:=1;i<=m;i++{
	fmt.Scan(&num)
		if num != (n*i) {
		benar = false
		break
		}
	}
	fmt.Println(benar)
}