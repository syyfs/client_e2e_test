package util

import "strings"

func FormatStateKey(args ...string) string {
	return strings.Join(args, "_")
}
