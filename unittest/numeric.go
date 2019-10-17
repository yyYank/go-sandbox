package main

import "strconv"

func isNumeric(value string) bool {
	if _, err := strconv.Atoi(value); err != nil {
		return false
	}
	return true
}
