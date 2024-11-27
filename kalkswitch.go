package main

import "fmt"

func main () {
	var x, y float64
	var op string
	
	fmt.Scan(&x, &op, &y)
	
	switch op {
	case "+" :
		fmt.Printf("%.2f\n", x + y)
	case "-" :
		fmt.Printf("%.2f\n", x - y)
	case "*" :
		fmt.Printf("%.2f\n", x * y)
	case "/" :
		fmt.Printf("%.2f\n", x / y)
	default :
		fmt.Println("lol dek")
	}
}