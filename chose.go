package main

import "fmt"

const g = 9.8

func main() {
	var n, i int
	var v, hasil float64
	fmt.Scan(&n, &v)
	
	for i = 0; i <= n; i++ {
	hasil = v*float64(i) + (0.5 * g * float64(i) * float64(i))
	fmt.Printf("%d %.5f \n",i, hasil)
	}

}
