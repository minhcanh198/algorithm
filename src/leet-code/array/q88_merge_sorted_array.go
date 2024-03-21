package array

// Merge
//Example 1:
//
//Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//Output: [1,2,2,3,5,6]
//Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
//The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
//Example 2:
//
//Input: nums1 = [1], m = 1, nums2 = [], n = 0
//Output: [1]
//Explanation: The arrays we are merging are [1] and [].
//The result of the merge is [1].
//Example 3:
//
//Input: nums1 = [0], m = 0, nums2 = [1], n = 1
//Output: [1]
//Explanation: The arrays we are merging are [] and [1].
//The result of the merge is [1].
//Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

func Merge(nums1 []int, m int, nums2 []int, n int) {
	leftArr := nums1[:m]
	i := 0
	j := 0

	mergedArr := []int{}
	for i < m && j < n {
		if leftArr[i] < nums2[j] {
			mergedArr = append(mergedArr, leftArr[i])
			i++
		} else {
			mergedArr = append(mergedArr, nums2[j])
			j++
		}
	}

	for j < n {
		mergedArr = append(mergedArr, nums2[j])
		j++
	}
	for i < m {
		mergedArr = append(mergedArr, leftArr[i])
		i++
	}

	k := 0
	for k < m+n {
		nums1[k] = mergedArr[k]
		k++
	}
}
