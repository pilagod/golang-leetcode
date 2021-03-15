package sort

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite
}

func (s *Suite) TestSort() {
	for _, t := range []struct {
		nums []int
	}{
		{
			[]int{1},
		},
		{
			[]int{2, 1},
		},
		{
			[]int{2, 3, 1},
		},
		{
			[]int{5, 4, 3, 2, 1},
		},
		{
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{5, 1, 4, 2, 3},
		},
		{
			[]int{2, 4, 5, 3, 1},
		},
		{
			[]int{5, 3, 4, 1, 2},
		},
		{
			[]int{6, 1, 5, 2, 4, 3},
		},
		{
			[]int{1, 3, 5, 6, 4, 2},
		},
		{
			[]int{6, 4, 5, 1, 3, 2},
		},
	} {
		input := make([]int, len(t.nums))
		copy(input, t.nums)
		quickSortResult := quickSort(input)
		s.True(check(quickSortResult, len(t.nums)), "quick sort %v fails: %v", t.nums, quickSortResult)
		s.Equal(t.nums, input)
	}
}

func check(nums []int, length int) bool {
	if len(nums) != length {
		return false
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
