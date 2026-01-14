package helpers

//check if message has only printable characters

func isPrintable(message string) bool {
	for _, char := range message {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
