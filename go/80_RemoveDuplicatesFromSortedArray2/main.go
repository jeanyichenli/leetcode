package main

func removeDuplicates2(nums []int) int {
	k := 0
	set := make(map[int]int)
	maxCount := 2

	for _, n := range nums {
		v, ok := set[n]
		if !ok || (ok && v < maxCount) { // none or first time
			set[n]++
			nums[k] = n
			k++
		}
	}

	return k
}
