package main

import "fmt"

const count = 15

func main() {
	i := 0
	for i < 10 {
		fmt.Print("LOOP :", i, "  ")
		i += 1
	}
	var input string
	for {
		fmt.Scanf("%s", &input)
		fmt.Println("Input ", input)
		if input == "exits" {
			break
		}
	}
}
