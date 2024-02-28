package sorts

import "fmt"

func MergeSort(arr []int, l int, r int) {
	fmt.Println("input", arr)
	if r == l {
		return
	}

	if r-l == 1 {
		if arr[r] <= arr[l] {
			swap(arr, r, l)
		}
		return
	}



}
