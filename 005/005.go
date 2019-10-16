package main

func longestPalindrome(s string) string {
	var result string
	for i := range s {
		var got string
		e1 := expand(s, i, i)
		e2 := expand(s, i, i+1)
		if len(e1) > len(e2) {
			got = e1
		} else {
			got = e2
		}
		if len(got) > len(result) {
			result = got
		}
	}
	return result
}

func expand(s string, i1, i2 int) string {
	if i1 < 0 || i2 >= len(s) || s[i1] != s[i2] {
		return ""
	}
	l, r := i1-1, i2+1
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
