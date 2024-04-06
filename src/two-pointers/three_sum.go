package two_pointers

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/3sum/?envType=study-plan-v2&envId=top-interview-150
func ThreeSum(nums []int) [][]int {
	fmt.Println(nums)
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1
		for l < r {
			if l == i {
				l++
				continue
			}
			if r == i {
				r--
				continue
			}
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
			}
		}
	}
	fmt.Println(res)
	return res
}
