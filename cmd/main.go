package main

import (
	"fmt"
	"github.com/minhcanh198/algorithm/src/sorts"
)

func main() {
	fmt.Println("mica course")

	arr := []int{2,4,5,10,10,21,1,2,12}

	//sorts.QuickSort(arr, 0, len(arr)-1)
	sorts.QuickSortPivotCenter(arr, 0, len(arr)-1)
	fmt.Println("res", arr)
}
