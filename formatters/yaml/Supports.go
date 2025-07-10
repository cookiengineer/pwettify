package yaml

import "gopkg.in/yaml.v3"

func Supports(buffer []byte) bool {

	var result bool
	var data interface{}

	err := yaml.Unmarshal(buffer, &data)

	if err == nil {
		result = true
	}

	return result

}
