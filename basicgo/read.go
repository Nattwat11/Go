package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("IndexInfo.csv")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		item := strings.Split(line, ",")

		//fmt.Println("Line %d : %S\n", line)
		fmt.Println(item[0])
		count++
	}
}
