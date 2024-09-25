package main

import (
	"strings"
)

func countSubstrings(s string) int {
	s = insertSymbolsToString(s)
	lenStr := len(s)
	P := make([]int, lenStr)

	C, R := 0, 0

	for i := 1; i < lenStr-1; i++ {
		i_mirror := 2*C - i // C - (i-C)

		if i < R && P[i_mirror]+i < R {
			P[i] = P[i_mirror]
			continue
		} else if i < R && P[i_mirror]+i > R {
			P[i] = R - i
			for s[i-1-P[i]] == s[i+1+P[i]] {
				P[i]++
				R = i + P[i]
				C = i
			}
		} else {
			P[i] = 0
			for s[i-1-P[i]] == s[i+1+P[i]] {
				P[i]++
				R = i + P[i]
				C = i
			}
		}
	}

	count := 0
	for _, pval := range P {
		count += (pval + 1) / 2
	}

	return count
}

func insertSymbolsToString(s string) string {
	return "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
}
