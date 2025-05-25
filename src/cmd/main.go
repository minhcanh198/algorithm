package main

import (
	"fmt"
	"github.com/minhcanh198/algorithm/medium"
)

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	r := medium.LongestConsecutive(nums)

	fmt.Println(r)
}
