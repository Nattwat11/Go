package main

import "fmt"

func hello() {
	fmt.Println("Hello_Born_To_Dev")
}
func plus(value1 int, value2 int) int {
	return value1 + value2

}
func plus3value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}
func main() {
	hello()
	reslt := plus(1, 2)
	fmt.Println(reslt)
	result2 := plus3value(2, 3, -5)
	fmt.Println(result2)
}
