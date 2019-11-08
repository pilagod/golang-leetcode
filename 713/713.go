package main

func numSubarrayProductLessThanK(nums []int, k int) (result int) {
	product, anchor := 1, 0
	for i, num := range nums {
		product *= num
		for anchor <= i && product >= k {
			product /= nums[anchor]
			anchor++
		}
		result += i - anchor + 1
	}
	return
}
