package core_utils

func GetStringLen(s string) int {
	return len([]rune(s))
}