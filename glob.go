package glob

import "unicode/utf8"

// IsFullMatch returns true if value is full match with given pattern.
func IsFullMatch(value, pattern string) bool {
	if pattern == "" {
		return value == ""
	}

	m, n := utf8.RuneCountInString(value), utf8.RuneCountInString(pattern)
	i, j := 0, 0
	backI, backJ := -1, -1
	for i < m {
		if j < n && pattern[j] == '*' {
			backI = i + 1
			backJ = j + 1
			j++
		} else if j < n && (value[i] == pattern[j] || pattern[j] == '?') {
			i++
			j++
		} else if backI != -1 {
			i = backI
			j = backJ
			backI++
		} else {
			return false
		}
	}

	for j < n && pattern[j] == '*' {
		j++
	}

	return j >= n
}
