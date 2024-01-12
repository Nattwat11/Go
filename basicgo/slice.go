package main

import "fmt"

func main() {
	var coursename []string
	coursename = []string{"JAVA", "Python"}

	coursename = append(coursename, "Go", "C#", "C+")
	fmt.Println(coursename)
	courseweb := coursename[3:5]
	fmt.Println(courseweb)
}
