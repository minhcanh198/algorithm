package dynamic_program

// https://leetcode.com/problems/climbing-stairs/?envType=study-plan-v2&envId=top-interview-150
func ClimbStairs(n int) int {
	if n < 3 {
		return n
	}

	ways := make([]int, n+1)
	ways[0] = 1
	ways[1] = 1
	ways[2] = 2

	for i := 3; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}

	return ways[n]
}
