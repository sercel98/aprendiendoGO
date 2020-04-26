package main

import (
	"fmt"
	"math"
)

func main() {

	s := "aab"
	n := 882787

	respuesta := repeatedString(s, int64(n))
	fmt.Println(respuesta)

}

func repeatedString(s string, n int64) int64 {

	cantA := 0
	for _, char := range s {
		if char == 'a' {
			cantA++
		}
	}

	if cantA == 0 {
		return 0
	}

	proportionInt := math.Floor(float64(n) / float64(len(s)))

	proportionFloat := float64(n)/float64(len(s)) - proportionInt

	result := int64(proportionInt) * int64(cantA)

	i := math.Round(proportionFloat * float64(len(s)))

	for j := 0; j < int(i); j++ {
		if s[j] == 'a' {
			result++
		}
	}

	return int64(result)

}
