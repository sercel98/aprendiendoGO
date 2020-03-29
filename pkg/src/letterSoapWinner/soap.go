package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	loadLetterSoap(text)
}

func loadLetterSoap(text string) {

	fmt.Println(len(text))
	matrix := [28][28]string{}

	textIndex := 0

	for countRow := 0; countRow < 28; countRow++ {
		for countCol := 0; countCol < 28 && textIndex <= len(text); countCol++ {
			matrix[countRow][countCol] = string(text[textIndex])

			fmt.Println(textIndex)

			textIndex++
		}
	}

	fmt.Println(matrix)
}
