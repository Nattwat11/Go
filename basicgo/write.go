package main

import (
	"os"
)

func main() {
	data1 := []byte("Hello \n BornTobe")
	err := os.WriteFile("DATA.txt", data1, 0644)
	if err != nil {
		panic(err)
	}
	f, ferr := os.Create("Employeename")
	if ferr != nil {
		panic(ferr)
	}

	defer f.Close()
	data2 := []byte("Sira \n Amanee")
	os.WriteFile("employeeName.txt", data2, 0644)
}
