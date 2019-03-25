package Truncate_String

func Truncate(s string, l int) string {
	if len(s) > l && l > 3 {
		return s[:l-3] + "..."
	}
	return s
}
