package math

// https://leetcode.com/problems/powx-n/?envType=study-plan-v2&envId=top-interview-150
func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	return x * MyPow(x, n-1)
}
