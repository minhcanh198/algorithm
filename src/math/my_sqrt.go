package math

// https://leetcode.com/problems/sqrtx/?envType=study-plan-v2&envId=top-interview-150
func MySqrt(x int) int {
	if x <= 1 {
		return x
	}

	for i := 1; i <= x/2+1; i++ {
		if i*i == x {
			return i
		}
		if i*i > x {
			return i - 1
		}
	}
	return 0
}
