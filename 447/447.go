package main

func numberOfBoomerangs(ps [][]int) int {
	ds := make(map[int](map[int]int))
	for i := 0; i < len(ps); i++ {
		ds[i] = make(map[int]int)
		for j := 0; j < len(ps); j++ {
			if i == j {
				continue
			}
			ds[i][distance(ps[i], ps[j])]++
		}
	}
	result := 0
	for i := range ds {
		for j := range ds[i] {
			result += c2(ds[i][j]) * 2
		}
	}
	return result
}

func distance(p1, p2 []int) int {
	dx, dy := p1[0]-p2[0], p1[1]-p2[1]
	return dx*dx + dy*dy
}

func c2(total int) int {
	return (total * (total - 1)) / 2
}
