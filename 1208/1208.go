package main

func equalSubstring(s string, t string, maxCost int) (result int) {
	if s == t {
		return len(s)
	}
	diffs := make([]int, len(s))
	for i := range s {
		diff := int(s[i]) - int(t[i])
		if diff < 0 {
			diff = -diff
		}
		diffs[i] = diff
	}
	sum, count, anchor := 0, 0, 0
	for i, diff := range diffs {
		for sum+diff > maxCost && anchor < i {
			sum -= diffs[anchor]
			count--
			anchor++
		}
		if anchor == i && diff > maxCost {
			anchor = i + 1
		} else {
			sum += diff
			count++
		}
		if count > result {
			result = count
		}
	}
	return
}
