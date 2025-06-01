package easy

// https://leetcode.com/problems/find-common-characters/
func CommonChars(words []string) []string {
	var commonChars []string
	wordCharMap := make([]map[uint8]int, len(words))

	for i, word := range words {
		wordCharMap[i] = make(map[uint8]int)
		for _, char := range word {
			wordCharMap[i][uint8(char)]++
		}
	}

	firstWordChar := wordCharMap[0]
	for char, charCount := range firstWordChar {
		minCount := charCount
		exist := 1
		for i, m := range wordCharMap {
			if i == 0 {
				continue
			}

			if count, ok := m[char]; ok {
				minCount = MinInt(minCount, count)
				exist++
			}
		}
		if exist == len(words) {
			for i := 0; i < minCount; i++ {
				commonChars = append(commonChars, string(char))
			}
		}
	}

	return commonChars
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
