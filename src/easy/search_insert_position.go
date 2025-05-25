package easy

// https://leetcode.com/problems/search-insert-position/?envType=study-plan-v2&envId=top-interview-150
func SearchInsertPosition(nums []int, target int) int {
	return BinarySearchInsertPosition(nums, target, 0, len(nums))
}

func BinarySearchInsertPosition(nums []int, find, left, right int) int {
	if left >= right {
		return left
	}

	mid := (left + right) / 2
	if nums[mid] == find {
		return mid
	}

	if nums[mid] < find {
		return BinarySearchInsertPosition(nums, find, mid+1, right)
	} else {
		return BinarySearchInsertPosition(nums, find, left, mid)
	}
}
