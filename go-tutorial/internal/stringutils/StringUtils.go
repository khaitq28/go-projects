package stringutils

import "strings"

func Concat(a, b string) string {
	return a + b + " from internal"
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func FindString(s string, substr string) bool {
	return strings.Contains(s, substr)
}

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[rune]int)
	for _, c := range s1 {
		m[c]++
	}
	for _, c := range s2 {
		m[c]--
		if m[c] < 0 {
			return false
		}
	}

	return true
}

func IsPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func CountTypes(s string) (lower, upper, digit int) {
	lo, hi, di := 0, 0, 0
	for _, r := range s {
		switch {
		case r >= '0' && r <= '9':
			di++
		case r >= 'a' && r <= 'z':
			lo++
		case r >= 'A' && r <= 'Z':
			hi++
		}
	}
	return lo, hi, di
}

func IsNumeric(s string) bool {
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}
