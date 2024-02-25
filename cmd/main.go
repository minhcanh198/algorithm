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
	inputStr := "minhcanh"
	fmt.Println(inputStr)
	log.Println(sliding_window.LengthOfLongestSubstring(inputStr))

	log.Println("sliding_window.MinSubArrayLen:", sliding_window.MinSubArrayLen(7, []int{2,3,1,2,4,3}))
	log.Println("sliding_window.MinSubArrayLen:", sliding_window.MinSubArrayLen(4, []int{1,4,4}))
	log.Println("sliding_window.MinSubArrayLen:", sliding_window.MinSubArrayLen(11, []int{1,1,1,1,1,1,1,1}))
}
