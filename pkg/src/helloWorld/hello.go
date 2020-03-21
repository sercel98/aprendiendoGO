package main

import "fmt"

func main() {

	serie := make([]int, 0)

	var nextNumber int = 1

	for nextNumber < 1000 {

		serie = append(serie, nextNumber)
		nextNumber++
	}

	fmt.Println(serie)
}
