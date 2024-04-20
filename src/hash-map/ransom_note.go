package hash_map

//https://leetcode.com/problems/ransom-note/?envType=study-plan-v2&envId=top-interview-150

func RansomCanConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int, 26)
	for _, letter := range magazine {
		magazineMap[letter] += 1

	}

	for _, letter := range ransomNote {
		magazineMap[letter] -= 1
		if magazineMap[letter] < 0 {
			return false
		}
	}

	return true
}
