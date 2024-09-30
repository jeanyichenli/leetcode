package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)

	result := [][]int{}

	for i := 0; i < n-2; i++ { // Fix one point and move other two pointers
		if i > 0 && nums[i] == nums[i-1] { // pass duplicate, i moves to right, so need to check the left one
			continue
		}

		left := i + 1
		right := n - 1

		for left < right {
			if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else { // sum == 0
				result = append(result, []int{nums[i], nums[left], nums[right]})
				right--
				left++

				for left < right && nums[left] == nums[left-1] { // pass duplicate, left moves right, so need to check left one
					left++
				}

				for left < right && nums[right] == nums[right+1] { // pass duplicate, right moves to left, so need to check right one
					right--
				}
			}
		}
	}

	return result
}
