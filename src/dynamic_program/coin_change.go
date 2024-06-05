package dynamic_program

// https://leetcode.com/problems/coin-change/?envType=study-plan-v2&envId=top-interview-150
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				if i-coin != 0 && dp[i-coin] == 0 {
					dp[i] = -1
					continue
				}

				if dp[i-coin] == -1 {
					continue
				}

				totalCoin := 1 + dp[i-coin]
				if dp[i] == 0 || dp[i] > totalCoin {
					dp[i] = 1 + dp[i-coin]
				}
			}
		}

		if dp[i] == 0 {
			dp[i] = -1
		}
	}

	return dp[amount]
}

//if i-coin >= 0 {
//				if dp[i-coin]+1 < dp[i] || dp[i] == 0 {
//					dp[i] = dp[i-coin] + 1
//				}
//			}
