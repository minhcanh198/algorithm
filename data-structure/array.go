package data_structure

import (
	"math/rand"
	"time"
)

func NewRandomArray(size int) []int {
	var arr []int
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(10000-100) + 100
		arr = append(arr, number)
	}
	return arr
}
