package sorts

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
	QuickSort(arr, newPivot + 1, r)
}
