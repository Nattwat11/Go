package main

import "fmt"

var productname [4]string

func main() {
	productname[0] = "Apple"
	productname[1] = "Samsung"
	productname[2] = "Vivo"
	productname[3] = "Xiaomi"
	price := [4]float32{25000, 30000, 7000, 3000}
	fmt.Println(productname)
	fmt.Println(price)
}
