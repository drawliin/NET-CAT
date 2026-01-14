package helpers

// check if name has only alphanumerical characters, dash and underscore
func isValidName(name string) bool {
	for _, char := range name {
		if !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') && !(char >= '0' && char <= '9') && !(char == '-' || char == '_') {
			return false
		}
	}
	return true
}
