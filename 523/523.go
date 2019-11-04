package main

func checkSubarraySum(nums []int, k int) bool {
	sums := make([]int, len(nums)+1)
	sums[0] = 0
	for i, n := range nums {
		sums[i+1] = sums[i] + n
	}
	dict := make(map[int]bool)
	dict[0] = true
	for i := 2; i < len(sums); i++ {
		if k == 0 && dict[sums[i]] || k != 0 && dict[sums[i]%k] {
			return true
		}
		if k == 0 {
			dict[sums[i-1]] = true
		} else {
			dict[sums[i-1]%k] = true
		}
	}
	return false
}

func checkSubarraySumBrute(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k || (k != 0 && sum%k == 0) {
				return true
			}
		}
	}
	return false
}
