package main

// Write a function that takes an array of integers and returns the largest subarray sum.
// we should be able to handle arrays containing negative elements
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
