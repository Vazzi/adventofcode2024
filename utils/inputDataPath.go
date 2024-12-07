package utils

import (
	"fmt"
)

func InputDataPath(day int, test bool) string {
	fileName := fmt.Sprintf("day%02d.txt", day)

	if test {
		return "../inputData/test/" + fileName
	}

	return "./inputData/" + fileName
}
