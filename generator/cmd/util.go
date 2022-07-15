package cmd

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FirstLower(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func Title(s string) string {
	return cases.Title(language.English).String(s[:1]) + s[1:]
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
