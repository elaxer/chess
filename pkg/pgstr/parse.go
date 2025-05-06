package pgstr

import (
	"strings"
)

func Parse(pgArray string) []string {
	str := strings.Trim(pgArray, "{}")
	elements := strings.Split(str, ",")
	for i := range elements {
		elements[i] = strings.Trim(elements[i], `"`)
	}

	return elements
}
