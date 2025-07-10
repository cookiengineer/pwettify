package json

import "encoding/json"

func Supports(buffer []byte) bool {

	var result bool
	var data interface{}

	if err := json.Unmarshal(buffer, &data); err == nil {
		result = true
	}

	return result

}
