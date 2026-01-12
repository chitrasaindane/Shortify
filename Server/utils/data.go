package utils

import "strings"

// # Get 'Trimmed Value'
func GetTrimmedValue(value string) string {
	trimmedValue := strings.TrimSpace(value)
	return trimmedValue
}

// # Get 'String Value'
func GetStringValue(value *string) string {
	var valueStr string
	if value != nil {
		valueStr = GetTrimmedValue(*value)
	} else {
		valueStr = ""
	}
	return valueStr
}
