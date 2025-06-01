package hard

// https://leetcode.com/problems/minimum-window-substring/description/?envType=study-plan-v2&envId=top-interview-150
type res struct {
	startIdx int
	endIdx   int
}

func MinWindow(s string, t string) string {
	if s == t {
		return s
	}
	if len(t) > len(s) {
		return ""
	}

	// Build frequency map for chars in t
	tmap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tmap[t[i]]++
	}

	smap := make(map[byte]int)
	l, r := 0, 0
	charOk := 0 // number of chars that meet required count
	minSub := res{startIdx: 0, endIdx: len(s) + 1}
	found := false

	for r < len(s) {
		// Expand window by including s[r]
		c := s[r]
		smap[c]++

		if count, ok := tmap[c]; ok && smap[c] == count {
			charOk++
		}

		// Shrink window from left while all chars matched
		for charOk == len(tmap) && l <= r {
			found = true

			// Update min window if smaller
			if (r - l + 1) < (minSub.endIdx - minSub.startIdx) {
				minSub = res{startIdx: l, endIdx: r + 1}
			}

			leftChar := s[l]
			smap[leftChar]--

			if count, ok := tmap[leftChar]; ok && smap[leftChar] < count {
				charOk--
			}

			l++
		}

		r++
	}

	if !found {
		return ""
	}
	return s[minSub.startIdx:minSub.endIdx]
}
