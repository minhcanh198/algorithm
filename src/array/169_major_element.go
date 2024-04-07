package array

import "fmt"

// https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
//func MajorityElement(nums []int) int {
//	minAppears := len(nums) / 2
//	appearMap := make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		appear, ok := appearMap[nums[i]]
//		if ok {
//			appear++
//			if appear > minAppears {
//				return nums[i]
//			}
//			appearMap[nums[i]] = appear
//		} else {
//			appearMap[nums[i]] = 1
//		}
//	}
//	return nums[0]
//}

func MajorityElement(nums []int) int {
	var ans int
	var count int

	for i, num := range nums {
		fmt.Println(i)
		if count == 0 {
			ans = num
		}
		if num == ans {
			count++
		} else {
			count--
		}
	}

	return ans
}
