package utils

func isPrintable(message string) bool {
	for _, char := range message {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
