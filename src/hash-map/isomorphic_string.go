package hash_map

// https://leetcode.com/problems/isomorphic-strings/description/?envType=study-plan-v2&envId=top-interview-150
func IsIsomorphic(s string, t string) bool {
	rep := ""
	wordMap := make(map[rune]rune)
	existMap := make(map[rune]bool)
	for i := 0; i < len(s); i++ {
		char, ok := wordMap[rune(s[i])]
		if ok {
			rep += string(char)
		} else {
			if existMap[rune(t[i])] {
				return false
			}
			existMap[rune(t[i])] = true
			wordMap[rune(s[i])] = rune(t[i])
			rep += string(t[i])
		}
	}

	return rep == t
}
