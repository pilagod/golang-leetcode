package main

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	v := make([][]bool, len(grid))
	for i := range v {
		v[i] = make([]bool, len(grid))
		for j := range v[i] {
			if grid[i][j] == 1 {
				v[i][j] = true
			}
		}
	}
	n := len(v) - 1
	step := 0
	curQueue := make([][]int, 0)
	curQueue = append(curQueue, []int{-1, -1})
	nextQueue := make([][]int, 0)
	for len(curQueue) > 0 {
		for _, c := range curQueue {
			if c[0] == n && c[1] == n {
				return step
			}
			for l := -1; l < 2; l++ {
				for k := -1; k < 2; k++ {
					x := c[0] + l
					y := c[1] + k
					if x >= 0 && x <= n && y >= 0 && y <= n && !v[x][y] {
						v[x][y] = true
						nextQueue = append(nextQueue, []int{x, y})
					}
				}
			}
		}
		step++
		curQueue, nextQueue = nextQueue, curQueue[:0]
	}
	return -1
}
