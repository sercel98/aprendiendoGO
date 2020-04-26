package main

import "fmt"

func main() {

	c := []int32{0, 0, 0, 1, 0, 0}

	result := jumpingOnClouds(c)
	fmt.Println(result)

}
func jumpingOnClouds(c []int32) int32 {

	steps := 0

	i := 0

	for i < len(c)-1 {

		if i+2 < len(c) && c[i+2] == 0 {
			i += 2
		} else {
			i++
		}

		steps++
	}
	return int32(steps)
}
