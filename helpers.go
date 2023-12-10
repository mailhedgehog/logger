package logger

// truncateStart
func truncateStart(str string, length int, omission string) string {
	r := []rune(str)
	sLen := len(r)
	if length >= sLen {
		return str
	}

	return string(omission + string(r[len(r)-length:]))
}
