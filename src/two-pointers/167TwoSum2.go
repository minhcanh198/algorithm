package two_pointers

import "fmt"

//https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/?envType=study-plan-v2&envId=top-interview-150
func TwoSum(numbers []int, target int) []int {
	fmt.Println(numbers, target)
	var res []int
	l := 0
	r := len(numbers) - 1
	for l < r {
		if numbers[l]+numbers[r] == target {
			res = []int{l + 1, r + 1}
			return res
		}
		if numbers[l]+numbers[r] > target {
			r--
			continue
		}
		if numbers[l]+numbers[r] < target {
			l++
			continue
		}
	}
	return res
}
