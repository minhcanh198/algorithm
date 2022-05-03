package sort

import "fmt"

func QuickSort(arr []int, left int, right int) {
	if left > right {
		return
	}

	pivotIndex := partition(arr, left, right)
	fmt.Println("pivotIndex")
	fmt.Println(pivotIndex)

	QuickSort(arr, left, pivotIndex)
	QuickSort(arr, pivotIndex+1, right)
}

func partition(arr []int, left int, right int) int {
	pivot := (left + right) / 2
	fmt.Println("pivot in part")
	fmt.Println(left, pivot, right)
	pivotElement := arr[pivot]
	fmt.Println("pivotElement")
	fmt.Println(pivotElement)
	low := left
	high := right
	for low < high {
		for arr[low] < pivotElement {
			low++
		}
		for arr[high] > pivotElement {
			high--
		}
		if low < high {
			temp := arr[low]
			arr[low] = arr[high]
			arr[high] = temp
			low++
			high--
		}
	}
	fmt.Println("arr")
	fmt.Println(arr)
	return low
}
