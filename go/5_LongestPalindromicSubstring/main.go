package main

import (
	"strings"
)

func longestPalindrome(s string) string {
	// renew string
	str := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	lenStr := len(str)

	R, C := 0, 0 // pointers
	P := make([]int, lenStr)

	for i := 1; i < lenStr-1; i++ {
		i_mirror := C - (i - C)
		if i < R && P[i_mirror]+i > R {
			P[i] = R - i
			for str[i+P[i]+1] == str[i-P[i]-1] {
				P[i]++
				R = i + P[i]
				C = i
			}
		} else if i < R && P[i_mirror]+i < R {
			P[i] = P[i_mirror]
		} else {
			P[i] = 0
			for str[i+P[i]+1] == str[i-P[i]-1] {
				P[i]++
				R = i + P[i]
				C = i
			}
		}
	}

	maxLen := 0
	maxCenter := 0

	for i, pval := range P {
		if pval > maxLen {
			maxLen = pval
			maxCenter = i
		}
	}

	return s[(maxCenter-maxLen)/2 : (maxCenter+maxLen)/2]
}
