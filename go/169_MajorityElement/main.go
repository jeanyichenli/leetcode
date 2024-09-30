package main

func majorityElement(nums []int) int {
	set := make(map[int]int)
	count := len(nums) / 2

	for _, n := range nums {
		set[n]++
		if set[n] > count {
			return n
		}
	}

	return 0
}
