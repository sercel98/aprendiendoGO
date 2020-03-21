package main

import "fmt"

func main() {

	var n int = 20

	fmt.Println(getFactorial(n))

}

func getFactorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * getFactorial(n-1)
}
