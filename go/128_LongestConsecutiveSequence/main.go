package main

func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}

	result := 0
	for _, num := range nums {
		if set[num-1] {
			continue
		}

		count := 0
		for n := num; set[n]; n++ {
			count++
		}

		if count > result {
			result = count
		}
	}

	return result
}
