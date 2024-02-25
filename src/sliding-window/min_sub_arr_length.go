package sliding_window

import (
	"fmt"
	"math"
)

func MinSubArrayLen(target int, nums []int) int {
	min := 9999999999999
	right := 0
	left := 0
	sum := 0
	change := false
	for right < len(nums) && left <= right {
		//fmt.Println("right", right)
		fmt.Println("sum:", sum, "left", left, "right", right, "min", min)
		if sum == target {
			min = int(math.Min(float64(min), float64(right-left+1)))
			change = true
			sum -= nums[left]
			left++
			right++
			continue
		}
		if sum < target {
			sum += nums[right]
			right++
			continue
		}

		if sum > target {
			sum -= nums[left]
			left++
			right --
		}
	}
	if !change {
		return 0
	}
	return min
}
