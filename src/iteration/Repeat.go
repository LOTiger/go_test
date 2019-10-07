package iteration

const repeatCount = 5

// Repeat is function to repeat a string
func Repeat(character string) string {
	var repected string
	for index := 0; index < repeatCount; index++ {
		repected = repected + character
	}

	return repected
}