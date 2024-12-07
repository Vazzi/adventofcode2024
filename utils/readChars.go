package utils

func ReadChars(fileName string) [][]string {
	lines := ReadLines(fileName)

	data := make([][]string, 0)

	for _, line := range lines {
		runes := []rune(line)
		chars := make([]string, len(runes))
		for i := range runes {
			chars[i] = string(runes[i])
		}
		data = append(data, chars)
	}

	return data
}
