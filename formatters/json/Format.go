package json

import "encoding/json"

func Format(buffer []byte) []byte {

	result := make([]byte, 0)

	var data interface{}

	err1 := json.Unmarshal(buffer, &data);

	if err1 == nil {

		formatted, err2 := json.MarshalIndent(data, "", "\t")

		if err2 == nil {
			result = formatted
		}

	}

	return result

}
