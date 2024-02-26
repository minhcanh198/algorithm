package sorts

import "fmt"

func QuickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := r
	runIdx := l
	newPivot := runIdx
	for runIdx < r {
		if arr[runIdx] < arr[pivot] {
			swap(arr, newPivot, runIdx)
			newPivot++
		}
		runIdx++
	}

	swap(arr, newPivot, pivot)
	QuickSort(arr, l, newPivot-1)
	QuickSort(arr, newPivot+1, r)
}

func QuickSortPivotLeft(arr []int, l int, r int) {
	fmt.Println("input", arr)
	if l >= r {
		return
	}

	pivot := l
	runIdx := l + 1
	newPivot := l
	for runIdx <= r {
		if arr[runIdx] < arr[pivot] {

			newPivot++
			swap(arr, newPivot, runIdx)
		}
		runIdx++
	}


	swap(arr, newPivot, pivot)
	QuickSortPivotLeft(arr, l, newPivot-1)

	QuickSortPivotLeft(arr, newPivot+1, r)
}
