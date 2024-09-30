package main

func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0              // left and right pointer
	count := make([]int, 256) // slice used to count occurrence of specific character
	ans := 0

	for r < len(s) { // right pointer cannot over the string
		count[s[r]]++
		for count[s[r]] > 1 { // Same character occurrs, let left pointer goes to same position of right one.
			count[s[l]]-- // During the movement, reset the count array
			l++           // Move one position to right
		} // Until count[n] <= 1, right and left pointers are at the same and new position. Start next dynamic sliding window
		d := r - l + 1 // Calculate distance between right and left pointers
		if ans < d {
			ans = d
		}
		r++
	}
	return ans
}
