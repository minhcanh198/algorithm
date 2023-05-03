package sliding_window

import (
	"log"
	"math"
)

func LengthOfLongestSubstring(s string) int {
	log.Println("input string: ", s)
	maxLength := 0
	seen := make(map[string]int)
	for i := 0; i < len(s); i++ {
		_, ok := seen[string(s[i])]
		if !ok {
			seen[string(s[i])] = i
			maxLength = int(math.Max(float64(len(seen)), float64(maxLength)))
		} else {
			seen = make(map[string]int)
			seen[string(s[i])] = i
		}
	}
	return maxLength
}
