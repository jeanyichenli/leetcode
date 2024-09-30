package main

func isAnagram(s string, t string) bool {
	count := make([]int, 26)

	for _, c := range s {
		i := int(c - 'a')
		count[i]++
	}

	for _, c := range t {
		i := int(c - 'a')
		v := count[i]
		if v > 0 {
			count[i]--
		} else {
			count[i]++
		}
	}

	for _, v := range count {
		if v > 0 {
			return false
		}
	}

	return true
}
