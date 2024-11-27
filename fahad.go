package main

import "fmt"

func main() {
	var fahad int
	fmt.Print("Please rate how bitch Fahad is. Rate from 1 to 10: ")
	fmt.Scan(&fahad)

	if fahad > 0 && fahad < 5 {
		fmt.Println("Bitch, you are probably Fahad!")
	} else if fahad >= 5 && fahad <= 10 {
		fmt.Println("That's RIGHT! FAHAD IS A BITCH!")
	} else {
		fmt.Println("Please fill it right, you dumbass!")
	}
}
