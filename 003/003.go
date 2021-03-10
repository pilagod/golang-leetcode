package main

func lengthOfLongestSubstring(s string) int {
	chars := []rune(s)
	visited := make(map[rune]int)
	result, left := 0, 0
	for right := 0; right < len(chars); right++ {
		if _, ok := visited[chars[right]]; ok && left <= visited[chars[right]] {
			left = visited[chars[right]] + 1
		}
		visited[chars[right]] = right
		if right-left+1 > result {
			result = right - left + 1
		}
	}
	return result
}
