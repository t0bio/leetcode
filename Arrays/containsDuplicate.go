package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for i := range m {
		if m[i] > 1 {
			return true
		}
	}
	return false
}
