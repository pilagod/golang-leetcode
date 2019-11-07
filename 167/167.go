package main

func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l != r {
		sum := nums[l] + nums[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}
		if sum > target {
			r--
		} else {
			l++
		}
	}
	return nil
}
