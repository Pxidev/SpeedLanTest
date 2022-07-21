package utils

import "regexp"

func IsNumber(value string) bool {
	pattern := regexp.MustCompile("^[0-9]*$")
	matches := pattern.MatchString(value)

	return matches
}
