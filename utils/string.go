package utils

import "strings"

func CheckStringIsBlank(s string) bool {
	if len(s) <= 0 || s == "" || strings.Count(s, " ") == len(s) {
		return true
	}
	return false
}
