package main

type tempWithIndex struct {
	temp  int
	index int
}

func dailyTemperatures(T []int) []int {
	if len(T) == 0 {
		return nil
	}
	result := make([]int, len(T))
	var stack []tempWithIndex
	for i, t := range T {
		for len(stack) > 0 && t > stack[len(stack)-1].temp {
			last := stack[len(stack)-1]
			result[last.index] = i - last.index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, tempWithIndex{t, i})
	}
	return result
}
