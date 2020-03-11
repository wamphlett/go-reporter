package main

import "github.com/fatih/color"

func IsInSlice(s string, list []string) bool {
	for _, check := range list {
		if check == s {
			return true
		}
	}

	return false
}

func PrintError(e string) {
	color.Red(e)
}
