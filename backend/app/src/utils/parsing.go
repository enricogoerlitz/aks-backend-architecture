package utils

import "strconv"

func ParseStringToInt(value string, defaultVal int) int {
	if value == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(value)
	if err != nil {
		return defaultVal
	}
	return val
}
