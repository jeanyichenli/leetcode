package main

func climbStairs(n int) int {
	cur := 0
	pre, post := 0, 0

	for i := 1; i <= n; i++ {
		if i == 1 {
			cur = 1
		} else if i == 2 {
			cur = 2
		} else {
			cur = pre + post
		}

		post = pre
		pre = cur
	}

	return cur
}
