package html

import "strings"

func Supports(buffer []byte) bool {

	var result bool

	check := strings.TrimSpace(string(buffer))

	if strings.HasPrefix(check, "<!DOCTYPE html") || strings.HasPrefix(check, "<!DOCTYPE HTML") {
		result = true
	}

	return result

}
