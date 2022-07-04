package main

func countOdds(low int, high int) []int {
	var nums []int
	for i := low; i <= high; i++ {
		if i%2 != 0 {
			nums = append(nums, i)
		}
	}
	return nums
}
