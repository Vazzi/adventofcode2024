package utils

func Copy2DArray(data [][]string) [][]string {
	copied := make([][]string, len(data))

	for i := range data {
		copied[i] = make([]string, len(data[i]))
		copy(copied[i], data[i])
	}

	return copied
}
