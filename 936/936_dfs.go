package main

import "strings"

/* TimeLimitExeceeded */
func movesToStampDFS(stamp string, target string) []int {
	if len(stamp) > len(target) {
		return nil
	}
	if len(stamp) == len(target) {
		if stamp == target {
			return []int{0}
		}
		return nil
	}
	if stamp[0] != target[0] ||
		stamp[len(stamp)-1] != target[len(target)-1] {
		return nil
	}
	pos := make([]int, len(target)-len(stamp)+1)
	for i := range pos {
		pos[i] = i
	}
	return dfs(nil, pos, strings.Repeat("?", len(target)), stamp, target)
}

func dfs(result, pos []int, cur, stamp, target string) []int {
	for i, p := range pos {
		next := cur[:p] + stamp + cur[p+len(stamp):]
		if next == target {
			return append(result, p)
		}
		dfsResult := dfs(
			append(result, p),
			append(append([]int{}, pos[:i]...), pos[i+1:]...),
			next,
			stamp,
			target,
		)
		if len(dfsResult) > 0 {
			return dfsResult
		}
	}
	return nil
}
