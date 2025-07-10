package xml

import "bytes"
import "encoding/xml"

func Format(buffer []byte) []byte {

	var result bytes.Buffer

	decoder := xml.NewDecoder(bytes.NewReader(buffer))

	encoder := xml.NewEncoder(&result)
	encoder.Indent("", "\t")

	for {

		token, err1 := decoder.Token()

		if err1 != nil {

			if err1.Error() == "EOF" {
				break
			} else {
				// Do Nothing?
			}

		}

		err2 := encoder.EncodeToken(token)

		if err2 != nil {
			// Do Nothing?
			break
		}

	}

	err3 := encoder.Flush()

	if err3 != nil {
		// Do Nothing?
	}

	return result.Bytes()

}
