package dynamic_program

// https://leetcode.com/problems/longest-increasing-subsequence/description/?envType=study-plan-v2&envId=top-interview-150
func LengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}

	dp := make([]int, len(nums))
	maxNum := 1

	for i := range nums {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = maxInt(dp[i], dp[j]+1)
				maxNum = maxInt(maxNum, dp[i])
			}
		}
	}

	return maxNum
}
