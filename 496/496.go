package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return nil
	}
	if len(nums2) == 0 {
		return handleEmptyNums2(nums1)
	}
	dict := make(map[int]int)
	stack := make([]int, len(nums2))
	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			dict[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	for _, num := range stack {
		dict[num] = -1
	}
	result := make([]int, len(nums1))
	for i, num := range nums1 {
		if nextGreaterNum, ok := dict[num]; ok {
			result[i] = nextGreaterNum
		} else {
			result[i] = -1
		}
	}
	return result
}

func handleEmptyNums2(nums1 []int) []int {
	result := make([]int, len(nums1))
	for i := range result {
		result[i] = -1
	}
	return result
}
