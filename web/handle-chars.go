package web

import "strings"

func handleSpaces(arg string) string {
	arg = strings.ReplaceAll(arg, " ", "+")
	return arg
}

func hasWildcard(arg string) bool {
	return strings.Contains(arg, "*") || strings.Contains(arg, "?")
}