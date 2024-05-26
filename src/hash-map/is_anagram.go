package hash_map

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	wordMap := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		wordMap[rune(s[i])]++
	}
	for i := 0; i < len(t); i++ {
		wordMap[rune(t[i])]--
		if wordMap[rune(t[i])] < 0 {
			return false
		}
	}

	return true
}
