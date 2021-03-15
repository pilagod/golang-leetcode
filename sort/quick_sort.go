package sort

func quickSort(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	pl, pr, pivot := partition(nums)
	l := quickSort(pl)
	r := quickSort(pr)
	return append(append(l, pivot), r...)
}

func partition(nums []int) (left, right []int, pivot int) {
	result := make([]int, len(nums))
	copy(result, nums)
	pivot = result[0]
	head, tail := 1, len(result)-1
	for head < tail {
		for head < tail && result[head] < pivot {
			head++
		}
		for tail > head && result[tail] > pivot {
			tail--
		}
		result[head], result[tail] = result[tail], result[head]
	}
	if pivot > result[head] {
		result[0], result[head] = result[head], result[0]
	}
	return result[:head], result[head+1:], pivot
}
