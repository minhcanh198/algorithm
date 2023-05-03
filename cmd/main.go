package main

import (
	"fmt"
	sliding_window "git.cps.onl/cps/quote/src/sliding-window"
	"git.cps.onl/cps/quote/src/sorts"
	"log"
)

func main() {
	arr := []int{1, 6, 3, 623, 34, 5, 2, 345, 2342}
	log.Println(arr)
	sorts.InterchangeSort(arr)
	log.Println(arr)
	inputStr := "abcaa"
	fmt.Println(inputStr)
	log.Println(sliding_window.LengthOfLongestSubstring(inputStr))
}
