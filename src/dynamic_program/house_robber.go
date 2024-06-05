package dynamic_program

// https://leetcode.com/problems/house-robber/?envType=study-plan-v2&envId=top-interview-150
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return maxInt(nums[0], nums[1])
	}

	maxRob := 0
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxInt(dp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		expectAsset := maxInt(dp[i-1], dp[i-2]+nums[i])
		dp[i] = expectAsset
		if dp[i] > maxRob {
			maxRob = dp[i]
		}
	}

	return maxRob
}

func maxInt(i, j int) int {
	if i < j {
		return j
	}
	return i
}
