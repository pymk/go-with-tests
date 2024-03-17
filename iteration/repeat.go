package iteration

func Repeat(character string, rep int) string {
	var repeated string
	for i := 0; i < rep; i++ {
		repeated += character
	}
	return repeated
}
