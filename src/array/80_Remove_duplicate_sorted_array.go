package array

import "fmt"

func RemoveDuplicates80(nums []int) int {
	fmt.Println(nums)
	if len(nums) == 1 {
		return 1
	}
	occurrence := 0
	idex := 0
	for i, num := range nums {
		if i-1 >= 0 && num == nums[i-1] {
			occurrence++
		} else {
			occurrence = 1
		}

		if occurrence <= 2 {
			nums[idex] = num
			idex++
		}
	}
	//fmt.Println(idex)
	return idex
}
