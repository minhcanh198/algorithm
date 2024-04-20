package array

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150
func MaxProfitII(prices []int) int {
	totalMaxProfit := 0
	maxProfit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			totalMaxProfit += maxProfit
			maxProfit = 0

			continue
		}

		profit := prices[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
		} else {
			totalMaxProfit += maxProfit
			minPrice = prices[i]
			maxProfit = 0
		}
	}

	if maxProfit > 0 {
		totalMaxProfit += maxProfit
	}

	return totalMaxProfit
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
