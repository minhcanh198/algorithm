package medium

func WordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]struct{})
	for _, s2 := range wordDict {
		wordMap[s2] = struct{}{}
	}

	dp := make([]bool, len(s))
	firstChar := string(s[0])
	if _, ok := wordMap[firstChar]; ok {
		dp[0] = true
	}

	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] == true {
				remain := s[j+1 : i+1]
				if _, ok := wordMap[remain]; ok {
					dp[i] = true
					break
				}
			} else {
				word := s[0 : i+1]
				if _, ok := wordMap[word]; ok {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(s)-1]
}
