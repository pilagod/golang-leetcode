package main

import (
	"sort"
)

type intervalForSorter struct {
	start int
	end   int
	index int
}

type intervalSorter []intervalForSorter

func (s intervalSorter) Len() int {
	return len(s)
}

func (s intervalSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s intervalSorter) Less(i, j int) bool {
	return s[i].start < s[j].start
}

func findRightInterval(intervals [][]int) []int {
	if len(intervals) == 0 {
		return nil
	}
	sorter := make(intervalSorter, len(intervals))
	for i, interval := range intervals {
		sorter[i] = intervalForSorter{interval[0], interval[1], i}
	}
	sort.Sort(sorter)
	result := make([]int, len(intervals))
	result[sorter[len(sorter)-1].index] = -1
	for i := 0; i < len(intervals)-1; i++ {
		j := i + 1
		for ; j < len(intervals); j++ {
			if sorter[i].end <= sorter[j].start {
				result[sorter[i].index] = sorter[j].index
				break
			}
		}
		if j == len(intervals) {
			result[sorter[i].index] = -1
		}
	}
	return result
}
