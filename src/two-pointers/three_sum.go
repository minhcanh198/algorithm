package two_pointers

import "fmt"

//https://leetcode.com/problems/3sum/?envType=study-plan-v2&envId=top-interview-150
func ThreeSum(nums []int) [][]int {
	fmt.Println(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		dict := make(map[int]int)
		target := nums[i]
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			shouldFound, ok := dict[-target-nums[j]]
			fmt.Println("should found", shouldFound, ok)
			if ok {
				res = append(res, []int{i, j, shouldFound})
			} else {
				dict[nums[j]] = j
			}
		}
	}
	fmt.Println(res)
	return res
}
