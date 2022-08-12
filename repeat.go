package iteration

func Repeat(char string, len int) (outString string) {
	for i := 0; i < len; i++ {
		outString += char
	}
	return
}
