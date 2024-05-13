package web

import "strings"

func handleSpaces(arg string) string {
	arg = strings.ReplaceAll(arg, " ", "+")
	return arg
}
