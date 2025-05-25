package medium

// https://leetcode.com/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-interview-150
func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	longest := 1
	numMap := make(map[int]struct{})
	for _, num := range nums {
		numMap[num] = struct{}{}
	}

	for num := range numMap {
		if _, ok := numMap[num-1]; ok {
			continue
		}
		count := 0
		next := num
		found := true
		for found {
			count++
			next = next + 1
			_, found = numMap[next]
		}

		longest = MaxInt(longest, count)
	}

	return longest
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
