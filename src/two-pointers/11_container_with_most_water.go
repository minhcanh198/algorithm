package two_pointers

// https://leetcode.com/problems/container-with-most-water/?envType=study-plan-v2&envId=top-interview-150
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		minHeight := minInt(height[left], height[right])
		area := minHeight * (right - left)
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
