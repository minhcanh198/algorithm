package array

//https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/?envType=study-plan-v2&envId=top-interview-150

func FirstOccurrenceStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			for j := i; j < len(haystack); j++ {
				if haystack[j] == needle[j-i] {
					if j-i == len(needle)-1 {
						return i
					}
					continue
				} else {
					break
				}
			}
		}
	}

	return -1
}
