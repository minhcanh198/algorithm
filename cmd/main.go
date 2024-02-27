package main

import (
	"fmt"
	"github.com/minhcanh198/algorithm/src/sorts"
)
//
//func main() {
//	// Example usage
//	arr := []int{2,4,5,10,10,21,1,2,12}
//	fmt.Println("Before sorting:", arr)
//	QuickSort(arr)
//	fmt.Println("After sorting:", arr)
//}
//
//// QuickSort sorts an integer slice using the QuickSort algorithm
//func QuickSort(arr []int) {
//	fmt.Println("input", arr)
//	if len(arr) <= 1 {
//		return
//	}
//
//	pivot := arr[len(arr)/2]
//	left := 0
//	right := len(arr) - 1
//	fmt.Println("pivot", pivot)
//
//	for left <= right {
//		for arr[left] < pivot {
//			left++
//		}
//
//		for arr[right] > pivot {
//			right--
//		}
//
//		if left <= right {
//			arr[left], arr[right] = arr[right], arr[left]
//			left++
//			right--
//		}
//		fmt.Println("after first looop", arr, left, right)
//
//	}
//
//	QuickSort(arr[:right+1])
//	//QuickSort(arr[left:])
//}

func main() {
	fmt.Println("mica course")

	arr := []int{2,4,5,10,10,21,1,2,12}

	//sorts.QuickSort(arr, 0, len(arr)-1)
	sorts.QuickSortPivotCenter(arr, 0, len(arr)-1)
	fmt.Println("res", arr)
}
