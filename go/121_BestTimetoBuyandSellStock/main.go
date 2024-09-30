package main

func maxProfit(prices []int) int {
	left, right := 0, 1
	ans := 0

	for right < len(prices) {
		if prices[right] > prices[left] {
			ans = max(ans, prices[right]-prices[left])
			right++
		} else {
			left = right
			right = left + 1
		}
	}
	return ans
}
