package main

import (
	"fmt"
	"github.com/minhcanh198/algorithm/medium"
)

func main() {
	str := "cars"
	wordDict := []string{"car", "ca", "rs"}
	r := medium.WordBreak(str, wordDict)

	fmt.Println(r)
}
