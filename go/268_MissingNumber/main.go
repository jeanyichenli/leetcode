package main

func missingNumber(nums []int) int {
	ans := len(nums)

	for i, n := range nums {
		ans += (i - n)
	}

	return ans
}
