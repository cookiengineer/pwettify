package html

import "golang.org/x/net/html"
import "bytes"

func Format(buffer []byte) []byte {

	var result bytes.Buffer

	root, err1 := html.Parse(bytes.NewReader(buffer))

	if err1 == nil {
		renderIndent(&result, root, 0)
	}

	return result.Bytes()

}
