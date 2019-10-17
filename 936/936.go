package main

import (
	"strings"
)

// It must exist at least one substring in target fully matching stamp.
// For those most or fully matching substrings, we can stamp them as late as possible.
// (1) Find all fully matching substrings in target
// (2) Find other partially matching substrings in target
// (3) Repeat (2) when there is no any matching exists
// (4) Reverse stamp positions found from (1) - (3)
// (5) If all positions are stamped, result is found
// (6) Otherwise, no result is there
func movesToStamp(stamp string, target string) []int {
	t, s := []byte(target), []byte(stamp)
	result := make([]int, 0, len(target))
	for hasStamped := true; hasStamped; {
		hasStamped = false
		for i := 0; i < len(t)-len(s)+1; i++ {
			if !canStamp(s, t, i) {
				continue
			}
			hasStamped = true
			for j := i; j < i+len(s); j++ {
				t[j] = '?'
			}
			result = append(result, i)
		}
	}
	if isStamped(t, 0, len(t)) {
		reverse(result)
		return result
	}
	return nil
}

func canStamp(s, t []byte, pos int) bool {
	if isStamped(t, pos, pos+len(s)) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if t[pos+i] == '?' {
			continue
		}
		if t[pos+i] != s[i] {
			return false
		}
	}
	return true
}

func isStamped(t []byte, l, r int) bool {
	return string(t[l:r]) == strings.Repeat("?", r-l)
}

func reverse(s []int) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
