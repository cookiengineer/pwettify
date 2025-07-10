package xml

import "strings"

func Supports(buffer []byte) bool {

	var result bool

	check := strings.TrimSpace(string(buffer))

	if strings.HasPrefix(check, "<?xml version=\"1.0\"") {
		result = true
	}

	return result

}
