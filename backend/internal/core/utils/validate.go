package core_utils

// ValidateStringLen checks if the length of the string is within the specified range [min, max].
// Returns true if the string length is valid (within range), false otherwise.
// example:
//
//	s := "string"
//	ok := ValidateStringLen(s, 1, 10)
//	if !ok {
//	    println("string len is not between 1 and 10")
//	}
func ValidateStringLen(s string, min int, max int) bool {
	sLen := len([]rune(s))
	return sLen >= min && sLen <= max
}