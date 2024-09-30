package main

func removeDuplicates(nums []int) int {
	k := 0
	set := make(map[int]struct{}) // map to record unique value

	for _, n := range nums {
		if _, ok := set[n]; !ok { // not exist
			nums[k] = n
			set[n] = struct{}{}
			k++
		}
	}
	return k
}
