package main

func numMatchingSubseq(s string, ws []string) int {
	dict := make(map[rune][]int)
	for i, c := range s {
		dict[c] = append(dict[c], i)
	}
	result := 0
	for _, w := range ws {
		last := -1
		for _, c := range w {
			is, ok := dict[c]
			if !ok || last >= is[len(is)-1] {
				last = -1
				break
			}
			for _, i := range is {
				if i > last {
					last = i
					break
				}
			}
		}
		if last != -1 {
			result++
		}
	}
	return result
}
