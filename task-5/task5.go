package task5

import (
	"strings"
)

func task5(input string) string {

	input = strings.ReplaceAll(input, "1", "one")

	return input
}
