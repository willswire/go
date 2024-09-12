package iteration

func Iterate(character string) string {
	var repeatedCharacter string
	for i := 0; i < 5; i++ {
		repeatedCharacter += character
	}
	return repeatedCharacter
}
