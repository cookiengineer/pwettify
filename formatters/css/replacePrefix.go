package css

import "strings"

func replacePrefix(input string, search string, replace string) string {

	if search == "" {
		return input
	}

	count := 0
	tmp := input

	for strings.HasPrefix(tmp, search) {
		tmp = tmp[len(search):]
		count++
	}

	prefix := ""

	for c := 0; c < count; c++ {
		prefix += replace
	}

	return prefix + tmp

}
