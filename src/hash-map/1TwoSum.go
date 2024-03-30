package hash_map

import "fmt"

//https://leetcode.com/problems/two-sum/description/?envType=study-plan-v2&envId=top-interview-150
func TwoSum(nums []int, target int) []int {
	fmt.Println(nums, target)
	var res []int
	dict := make(map[int]int)
	for i, num := range nums {
		//fmt.Println(i, num)
		if exKey, ok := dict[target-num]; ok {
			res = []int{exKey, i}
			return res
		}
		dict[num] = i

	}

	return res
}
