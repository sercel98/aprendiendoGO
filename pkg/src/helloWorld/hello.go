package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please, input the first number")
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Now, input the second number")
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Please, input the operation symbol(+,-,*,/) ")
	scanner.Scan()
	operation := scanner.Text()

	if operation == "+" {
		fmt.Println(sum(a, b))
	} else if operation == "-" {
		fmt.Println(rest(a, b))
	} else if operation == "*" {
		fmt.Println(multiply(a, b))
	} else {
		fmt.Println(divide(a, b))
	}

}

func sum(a int, b int) int {
	return a + b
}

func rest(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) float64 {
	n1 := float64(a)
	n2 := float64(b)
	return n1 / n2
}
