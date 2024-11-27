package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		hasil := ""
		for j := 0; j < i; j++ {
			hasil += "*" 
		}
		fmt.Println(hasil)
	}
}
