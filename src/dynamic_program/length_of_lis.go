package dynamic_program

// https://leetcode.com/problems/longest-increasing-subsequence/description/?envType=study-plan-v2&envId=top-interview-150
func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}

	}
	maxNum := 0
	for i := range dp {
		maxNum = maxInt(maxNum, dp[i])
	}
	return maxNum
}

//
//func LengthOfLIS(nums []int) int {
//}
