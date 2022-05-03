package main

import (
	data_structure "algorithm/data-structure"
	"algorithm/sort"
	"fmt"
)

func main() {
	arr := data_structure.NewRandomArray(3)
	arr = []int{9865, 4663, 7182}
	fmt.Println("source arr:", arr)
	//sort.InterchangeSort(arr)
	//sort.MergeSort(arr, 0, len(arr)-1)

	sort.QuickSort(arr, 0, len(arr)-1)
	fmt.Println("sorted:", arr)
}
