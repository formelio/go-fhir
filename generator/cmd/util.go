package cmd

import (
	"strings"
)

func FirstLower(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func FirstUpper(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
