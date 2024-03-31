package array

//https://leetcode.com/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
func RotateArray(nums []int, k int) {
	//fmt.Println(nums)
	tmp := make([]int, 0, len(nums))
	k = k % len(nums)
	//fmt.Println(nums[len(nums)-k:])
	tmp = append(tmp, nums[len(nums)-k:]...)
	tmp = append(tmp, nums[:len(nums)-k]...)

	//tmp = append(tmp, nums[:len])

	copy(nums, tmp)
}
