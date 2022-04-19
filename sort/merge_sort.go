package sort

func MergeSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2

	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)
	merge(arr, left, right)
}

func merge(arr []int, left int, right int) {
	mid := (left + right) / 2
	i := left
	j := mid + 1
	var mergedArr []int
	for i <= mid && j <= right {
		if arr[j] < arr[i] {
			mergedArr = append(mergedArr, arr[j])
			j++
		} else {
			mergedArr = append(mergedArr, arr[i])
			i++
		}
	}

	for i <= mid {
		mergedArr = append(mergedArr, arr[i])
		i++
	}
	for j <= right {
		mergedArr = append(mergedArr, arr[j])
		j++
	}

	i = left
	mergedArrIndex := 0
	for i <= right && mergedArrIndex < len(mergedArr) {
		arr[i] = mergedArr[mergedArrIndex]
		i++
		mergedArrIndex++
	}
}
