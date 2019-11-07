package main

type numWithIndex struct {
	num   int
	index int
}

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := make([]int, len(nums))
	for i := range result {
		result[i] = -1
	}
	var stack []numWithIndex
	for i, num := range append(nums, nums...) {
		for len(stack) > 0 && num > stack[len(stack)-1].num {
			last := stack[len(stack)-1]
			if result[last.index] == -1 {
				result[last.index] = num
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, numWithIndex{
			num:   num,
			index: i % len(nums),
		})
	}
	return result
}
