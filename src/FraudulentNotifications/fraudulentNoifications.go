package main

import (
	"fmt"
	"sort"
)

func main() {

	var expenditure1 = []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}
	var d1 int32 = 5
	result := activityNotifications(expenditure1, d1)
	fmt.Println(result)

	var expenditure2 = []int32{1, 2, 3, 4, 4}
	var d2 int32 = 4
	result2 := activityNotifications(expenditure2, d2)
	fmt.Println(result2)
}

func activityNotifications(expenditure []int32, d int32) int32 {

	intd := int(d)
	var notifications int32 = 0

	for i := 0; i+intd < len(expenditure); i++ {

		data := expenditure[i : i+intd]
		median := calcMedian(data)

		if median*2 <= float32(expenditure[i+intd]) {
			notifications++
		}
	}

	return notifications
}

func calcMedian(data []int32) float32 {

	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	middle := len(data) / 2

	if len(data)%2 != 0 {
		return float32(data[int(middle)])
	}

	median := (float32(data[middle-1]) + float32(data[middle])) / 2
	return median
}
