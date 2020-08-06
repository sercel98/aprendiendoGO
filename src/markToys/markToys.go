package main

import (
	"fmt"
	"sort"
)

func main() {

	prices := []int32{1,2,3,4}
	var budget int32 = 7

	result := maximumToys(prices, budget)

	fmt.Println(result)

}


func maximumToys(prices []int32, k int32) int32 {

	sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })
	toys := 0
	var moneySpend int32 = 0

	for _, price := range prices {

		if moneySpend + price <= k {
			moneySpend+= price
			toys++
		}

		if moneySpend >= k {
			break
		}
	}

	return int32(toys)
}

