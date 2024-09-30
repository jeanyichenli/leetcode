package main

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for _, n := range nums {
		if _, ok := set[n]; ok {
			return true
		} else {
			set[n] = struct{}{}
		}
	}
	return false
}
