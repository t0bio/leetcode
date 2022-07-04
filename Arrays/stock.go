package main

import "fmt"

func maxProfit(prices []int) int {
	currentmin := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < currentmin {
			currentmin = prices[i]
		} else if prices[i]-currentmin > profit {
			profit = prices[i] - currentmin
		}
	}

	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
