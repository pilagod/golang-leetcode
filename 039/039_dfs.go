package main

import "sort"

func combinationSumDFS(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target, len(candidates)-1)
}

func dfs(candidates []int, target, index int) (result [][]int) {
	for i := index; i >= 0; i-- {
		if candidates[i] == target {
			result = append(result, []int{candidates[i]})
			continue
		}
		if candidates[i] > target {
			continue
		}
		temp := dfs(candidates, target-candidates[i], i)
		for _, c := range temp {
			result = append(result, append(c, candidates[i]))
		}
	}
	return result
}
