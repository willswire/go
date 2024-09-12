package iteration

// Repeat the provided character a set number of times
func Repeat(character string, repeatCount int) string {
	var repeatedCharacter string
	for i := 0; i < repeatCount; i++ {
		repeatedCharacter += character
	}
	return repeatedCharacter
}
