package iteration

// Repeat takes a character, and creates a string containing it the specified number of time.
func Repeat(character string, count int) (repeated string) {

	for i := 0; i < count; i++ {
		repeated += character
	}
	return
}
