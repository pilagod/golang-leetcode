package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	cs := make([][][]int, target+1)
	cs[0] = append(cs[0], []int{})
	for _, n := range candidates {
		for v := n; v <= target; v++ {
			for _, c := range cs[v-n] {
				cn := make([]int, len(c)+1)
				copy(cn[:len(c)], c)
				cn[len(c)] = n
				cs[v] = append(cs[v], cn)
			}
		}
	}
	return cs[target]
}
