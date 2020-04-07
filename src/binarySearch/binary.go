package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please, input a number betweet 1 and 10")
	scanner.Scan()
	n, error := strconv.Atoi(scanner.Text())

	if error != nil {
		fmt.Println("We had some trouble. Sorry!")
		return
	}
	n = 3

	found := false
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	init := 0
	end := len(numbers)
	iterations := 1

	for init <= end && !found {
		middle := (init + end) / 2

		fmt.Println("i:", init, "m:", middle, "e:", end)

		if numbers[middle] == n {
			found = true
		} else if numbers[middle] > n {
			end = middle - 1
		} else {
			init = middle + 1
		}
		iterations++

	}
	fmt.Println(iterations)

}
