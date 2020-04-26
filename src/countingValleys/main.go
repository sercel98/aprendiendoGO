package main

import "fmt"

func main() {

	//example values
	var n int32 = 12
	var s = "DDUUDDUDUUUD"

	val := countingValleys(n, s)
	fmt.Println(val)

}

func countingValleys(n int32, s string) int32 {

	count := 0
	valleys := 0

	for i := int32(0); i < n; i++ {
		if s[i] == 'U' {
			count++
		}
		if s[i] == 'D' {
			count--

			j := i + 1

			for count < 0 && j < n {

				if s[j] == 'D' {
					count--
				} else {
					count++
				}
				if count == 0 {
					valleys++
					i = j
					break
				}
				j++
			}
		}
	}
	return int32(valleys)
}
