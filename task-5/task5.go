package task5

import (
	"strings"
)

func task5(input string) string {
	return strings.ReplaceAll(input, "1", "one")
}

func task5Var2(input string) string {
	var res []byte
	for _, elem := range input {
		if elem == '1' {
			res = append(res, []byte("one")...)
		} else {
			res = append(res, byte(elem))
		}
	}
	return string(res)
}
