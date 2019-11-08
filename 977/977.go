package main

func sortedSquares(A []int) []int {
	for i := range A {
		A[i] = A[i] * A[i]
	}
	result, i, l, r := make([]int, len(A)), 0, 0, len(A)-1
	for ; l != r; i++ {
		if A[l] > A[r] {
			result[i] = A[l]
			l++
		} else {
			result[i] = A[r]
			r--
		}
	}
	result[i] = A[l]
	reverse(result)
	return result
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
