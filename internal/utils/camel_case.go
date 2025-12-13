package utils

import "strings"

func GetCamelCase(s string) string {
	firstLetter := s[0]
	remainingString := s[1:]

	return strings.ToLower(string(firstLetter)) + remainingString
}
