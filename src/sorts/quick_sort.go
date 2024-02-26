package sorts

func QuickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := r

	pivotValue := arr[pivot]
	runIdx := l
	rememberIdx := runIdx

	for runIdx < r {
		if arr[runIdx] < pivotValue {
			swap(arr, rememberIdx, runIdx)
			rememberIdx++
		}

		runIdx++
	}
	swap(arr, rememberIdx, pivot)
	pivot = rememberIdx

	QuickSort(arr, l, pivot-1)
	QuickSort(arr, pivot+1, r)
}
