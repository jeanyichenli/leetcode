package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := 0

	for left < right {
		if height[left] < height[right] {
			area = max(area, height[left]*(right-left))
			left++
			continue
		} else {
			area = max(area, height[right]*(right-left))
			right--
			continue
		}
	}

	return area
}
