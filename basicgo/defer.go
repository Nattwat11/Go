package main

import "fmt"

func add(varlue1, value2 float64) {
	result := varlue1 + value2
	fmt.Println("result :", result)
}
func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i =", i)
	}
}
func deferloop() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("j=", i)
	}
}
func main() {
	fmt.Println("Welcom to Calculator")
	defer fmt.Println("END")
	add(20, 30)
	defer add(15, 20)
	add(90, 40)
	loop()
	deferloop()
}
