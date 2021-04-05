package utils

import "strings"

func ReturnContains(search string, s string) string {
	if strings.Contains(search, s) {
		return search
	} else {
		return ""
	}
}
