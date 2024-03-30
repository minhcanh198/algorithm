package two_pointers

import "fmt"

//https://leetcode.com/problems/is-subsequence/description/?envType=study-plan-v2&envId=top-interview-150
func IsSubsequence(s string, t string) bool {
	fmt.Println(s, t)
	if len(s) == 0 {
		return true
	}
	sLeft := 0
	l := 0
	for l < len(t) {
		if t[l] != s[sLeft] {
			//fmt.Println(l)
			l++
		} else {
			if sLeft == len(s)-1 {
				return true
			}
			sLeft++
			l++
		}
	}
	return false
}
