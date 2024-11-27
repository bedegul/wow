package main

import "fmt"


func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a + b > c && a + c > b && b + c > a {
	fmt.Print(a, ",", b, ", dan", c, "segitiga? true")
	} else {
		fmt.Print(a, ",", b, ", dan", c, "segitiga? false")
		}
}
