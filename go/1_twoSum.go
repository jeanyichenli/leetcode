package main

func twoSum(nums []int, target int) []int {
	valMap := make(map[int]int)
	result := []int{}

	for i, x := range nums {
		idx, ok := valMap[target-x]
		if ok {
			result := append(result, idx, i)
			return result
		}
		valMap[x] = i
	}

	return result
}
