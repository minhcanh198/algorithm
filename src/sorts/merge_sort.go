package sorts

import (
	"fmt"
)

func MergeSort(arr []int, l int, r int) {
	fmt.Println("input", arr)
	if r == l {
		return
	}

	if r-l == 1 {
		if arr[r] < arr[l] {
			swap(arr, r, l)
		}
		return

	}

	pivot := (l + r) / 2
	MergeSort(arr, l, pivot)
	MergeSort(arr, pivot+1, r)
	copy(arr[l:r+1], MergeSortedArrays(arr[l:pivot+1], arr[pivot+1:r+1]))
}

func MergeSortedArrays(arr1 []int, arr2 []int) []int {
	var mergedArr []int
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			mergedArr = append(mergedArr, arr1[i])
			i++
		} else {
			mergedArr = append(mergedArr, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		mergedArr = append(mergedArr, arr1[i])
		i++
	}
	for j < len(arr2) {
		mergedArr = append(mergedArr, arr2[j])
		j++
	}

	fmt.Println("mergedArr", mergedArr)
	return mergedArr
}
