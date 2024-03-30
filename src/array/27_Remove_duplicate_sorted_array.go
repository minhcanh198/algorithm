package array

func RemoveDuplicates(nums []int) int {
	//fmt.Println(nums)
	i := 0
	for _, num := range nums {
		if num != nums[i] {
			i++
			nums[i] = num
		}
	}
	//fmt.Println(i)
	//fmt.Println(nums)
	return i + 1
}
