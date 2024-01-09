package main

import "fmt"

func zerovalue(ivalue int) {
	ivalue = 0
}
func zeropointer(ipointer *int) {
	*ipointer = 0
}

func main() {
	i := 1
	fmt.Println("I = ", i)
	zerovalue(i)
	fmt.Println("value", i)
	zeropointer(&i)
	fmt.Println("Address", i)

}
