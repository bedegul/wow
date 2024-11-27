package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a+b > c && a+c > b && b+c > a {
		fmt.Printf("%d, %d, dan %d dapat membentuk segitiga? true\n", a, b, c)
	} else {
		fmt.Printf("%d, %d, dan %d dapat membentuk segitiga? false\n", a, b, c)
	}
}
