package yaml

import "gopkg.in/yaml.v3"
import "bytes"

func Format(buffer []byte) []byte {

	var result bytes.Buffer
	var root yaml.Node

	decoder := yaml.NewDecoder(bytes.NewReader(buffer))

	err1 := decoder.Decode(&root)

	if err1 != nil {
		// Do Nothing?
	}

	encoder := yaml.NewEncoder(&result)
	encoder.SetIndent(4) // 4 spaces

	err2 := encoder.Encode(&root)

	if err2 != nil {
		// Do Nothing?
	}

	err3 := encoder.Close()

	if err3 != nil {
		// Do Nothing?
	}

	return result.Bytes()

}
