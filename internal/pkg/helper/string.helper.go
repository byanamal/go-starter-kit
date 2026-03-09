package helper

import "strings"

func String(v string) *string {
	return &v
}

func TruncateString(str string, length int) string {
	if len(str) <= length {
		return str
	}
	return str[:length] + "..."
}

func ToTitleCase(str string) string {
	return strings.Title(strings.ToLower(str))
}
