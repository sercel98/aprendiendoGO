package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type calculator struct {
	number1 int
	number2 int
}

func (calculator) sum() int {
	return calculatornumber1 + number2
}

func (calculator) rest(a int, b int) int {
	return a - b
}

func (calculator) multiply(a int, b int) int {
	return a * b
}

func (calculator) divide(a int, b int) float64 {
	n1 := float64(a)
	n2 := float64(b)
	return n1 / n2
}
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please, input the first number")
	scanner.Scan()
	a, error := strconv.Atoi(scanner.Text())

	fmt.Println("Now, input the second number")
	scanner.Scan()
	b, error := strconv.Atoi(scanner.Text())

	if error != nil {
		fmt.Println("I'm sorry, there was an error!")
	}

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
