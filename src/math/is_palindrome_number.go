package math

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original := x
	reverseNumber := 0
	for original > 0 {
		reverseNumber = reverseNumber*10 + original%10
		original = original / 10
	}

	return x == reverseNumber
}
