package sliding_window

func LengthOfLongestSubstring(s string) int {
	maxLength := 0
	start := 0
	end := 0
	charSet := make(map[byte]bool)

	for end < len(s) {
		if !charSet[s[end]] {
			charSet[s[end]] = true
			end++
			maxLength = max(maxLength, end-start)
		} else {
			delete(charSet, s[start])
			start++
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
