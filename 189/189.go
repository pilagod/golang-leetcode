package main

func rotate(nums []int, k int) {
	step := k % len(nums)
	reverse(nums)
	reverse(nums[:step])
	reverse(nums[step:])
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
