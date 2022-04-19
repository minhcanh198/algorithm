package main

import (
	data_structure "algorithm/data-structure"
	"algorithm/sort"
	"fmt"
)

func main() {
	arr := data_structure.NewRandomArray(11)
	fmt.Println("source arr:", arr)
	//sort.InterchangeSort(arr)
	sort.MergeSort(arr, 0, len(arr)-1)
	fmt.Println("sorted:", arr)
}
