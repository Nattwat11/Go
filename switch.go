package main

import "fmt"

var input string

func main() {
	i := 0
	for i < 10 {
		fmt.Scanf("%s", &input)
		switch input {
		case "Blue":
			fmt.Println("#0000F")
		case "Green":
			fmt.Println("#008000")
		case "Pink":
			fmt.Println("#FFC0CB")
		case "Yellow":
			fmt.Println("#FFFF00")
		default:
			fmt.Println("No have Data to Database")
		}
		i += 1
	}
}
