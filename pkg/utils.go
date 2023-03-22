package utils

func IsValidPrompt(prompt string) bool {
	if len(prompt) > 500 {
		return false
	}
	if len(prompt) == 0 {
		return false
	}
	return true
}
