package main

import (
	"fmt"
	"github.com/minhcanh198/algorithm/src/sorts"
)

func main() {
	fmt.Println("mica course ==>>>>>>>>>>>>>>>>>>>>")

	arr := []int{1, 12, 5, 26, 7, 14, 3, 7, 2}

	sorts.QuickSortPivotCenter(arr, 0, len(arr)-1)
	fmt.Println("res", arr)
}
