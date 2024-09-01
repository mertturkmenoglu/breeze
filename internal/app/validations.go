package app

import "strings"

func isValidEmail(s string) bool {
	if len(s) < 6 {
		return false
	}

	containsAtSymbol := strings.Contains(s, "@")

	if !containsAtSymbol {
		return false
	}

	idx := strings.Index(s, "@")

	if idx == 0 || idx == len(s)-1 {
		return false
	}

	return true
}

func isValidPassword(s string) bool {
	return len(s) >= 8
}

func isValidName(s string) bool {
	return len(s) >= 2
}

func isValidLoginPassword(s string) bool {
	return len(s) >= 1
}
