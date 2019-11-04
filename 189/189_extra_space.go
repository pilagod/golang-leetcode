package main

func rotateExtraSpace(nums []int, k int) {
	length := len(nums)
	step := k % length
	// copy must ensure dest has enough length
	temp := make([]int, step)
	copy(temp, nums[length-step:])
	for i, j := length-step-1, 1; i >= 0; i, j = i-1, j+1 {
		nums[length-j] = nums[i]
	}
	for i, t := range temp {
		nums[i] = t
	}
}

func rotateExtraSpaceMod(nums []int, k int) {
	temp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		temp[(i+k)%len(nums)] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = temp[i]
	}
}
