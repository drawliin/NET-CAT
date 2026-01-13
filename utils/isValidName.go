package utils

// check whether name has only alphanumerical characters, dash and underscore
func isValidName(message string) bool {
	for _, char := range message {
		if !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') && !(char >= '0' && char <= '9') && !(char == '-' || char == '_') {
			return false
		}
	}
	return true
}
