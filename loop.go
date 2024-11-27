package main

import "fmt"

func main () {
	var x, y, total int
	fmt.Scan(&y, &x)
	
	for i := x; i <= y; i++ {
	total += i
	fmt.Println(i)
}
	
	fmt.Print("total = ", total)
}