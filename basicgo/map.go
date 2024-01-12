package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("Product = ", product)
	///add
	product["Macbook"] = 40000
	product["Ipad"] = 25000
	product["Airpod"] = 7000
	fmt.Println(product)
	delete(product, "Macbook")
	///update
	product["Ipad"] = 30000
	fmt.Println(product)
	value1 := product["Ipad"]
	fmt.Println("value =", value1)

	courseName := map[string]string{"101": "JAVA", "102": "Python"}
	fmt.Println("Coursename = ", courseName)
}
