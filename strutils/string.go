package strutils

// Checks if string is empty ("").
func IsEmpty(s string) bool {
	return len(s) == 0
}

// Checks if a string is not empty ("").
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// Checks if a string is empty ("") or whitespace only.
func IsBlank(s string) bool {
	if len(s) == 0 {
		return true
	}
	runes := []rune(s)
	for _, r := range runes {
		if r != ' ' {
			return false
		}
	}
	return true
}

// Checks if a string is not empty (""), and not whitespace only.
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}
