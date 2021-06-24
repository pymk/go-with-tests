package iteration

func Repeat(character string) string {

	const repeatCount = 5

	var repeated string // Same as repeated := ""

	for i := 0; i < repeatCount; i++ {
		repeated += character // Same as repeated = repeated + character
	}
	return repeated
}