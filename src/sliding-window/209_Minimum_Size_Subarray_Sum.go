package sliding_window

// https://leetcode.com/problems/minimum-size-subarray-sum/description/?envType=study-plan-v2&envId=top-interview-150
func MinSubArrayLen(target int, nums []int) int {
	minSubLen := len(nums)
	isFound := false
	i := 0
	sum := 0

	for j := i; j < len(nums); j++ {
		sum += nums[j]

		for sum >= target {
			isFound = true
			minSubLen = minInt(minSubLen, j-i+1)
			sum -= nums[i]
			i++
		}
	}

	if isFound {
		return minSubLen
	}

	return 0
}

func minInt(x int, y int) int {
	if x < y {
		return x
	}

	return y
}
