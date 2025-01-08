package main

import "strings"

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	result := strings.Fields(output)
	return result
}
