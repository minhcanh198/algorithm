package sliding_window

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/?envType=study-plan-v2&envId=top-interview-150
func LengthOfLongestSubstring(s string) int {
	n := len(s)
	longest := 0
	left := 0
	right := 0
	charSet := make(map[string]bool)
	for right < n {
		if !charSet[string(s[right])] {
			charSet[string(s[right])] = true
			longest = maxInt(longest, right-left+1)
			right++
		} else {
			delete(charSet, string(s[left]))
			left++
		}
	}
	return longest
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
