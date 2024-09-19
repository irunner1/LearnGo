package iteration

const repeatCount int = 5

// Repeat character `repeatCount“ times.
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
