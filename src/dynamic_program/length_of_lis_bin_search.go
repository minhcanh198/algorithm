package dynamic_program

import "github.com/minhcanh198/algorithm/easy"

// https://leetcode.com/problems/longest-increasing-subsequence/description/?envType=study-plan-v2&envId=top-interview-150
func LengthOfLISBinSearch(nums []int) int {
	var tails []int
	for _, num := range nums {
		idx := easy.BinarySearchInsertPosition(tails, num, 0, len(tails))
		if idx == len(tails) {
			tails = append(tails, num)
		} else {
			tails[idx] = num
		}
	}

	return len(tails)
}

func BinarySearch(nums []int, find, left, right int) int {
	if left >= right {
		return -1
	}

	mid := (left + right) / 2
	if nums[mid] == find {
		return mid
	}

	if nums[mid] < find {
		return BinarySearch(nums, find, mid+1, right)
	} else {
		return BinarySearch(nums, find, left, mid)
	}
}
