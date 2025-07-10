package html

import "golang.org/x/net/html"
import "bytes"
import "fmt"

func renderIndent(result *bytes.Buffer, node *html.Node, depth int) {

	switch node.Type {
	case html.DocumentNode:

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			renderIndent(result, child, depth)
		}

	case html.DoctypeNode:

		result.WriteByte('\n')

		for i := 0; i < depth; i++ {
			result.WriteByte('\t')
		}

		result.WriteString("<!DOCTYPE " + node.Data + ">")

	case html.ElementNode:


		result.WriteByte('\n')

		for i := 0; i < depth; i++ {
			result.WriteByte('\t')
		}

		result.WriteString("<" + node.Data)

		for _, attribute := range node.Attr {
			result.WriteString(fmt.Sprintf(` %s="%s"`, attribute.Key, attribute.Val))
		}

		result.WriteString(">")

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			renderIndent(result, child, depth + 1)
		}

		if node.FirstChild != nil {

			result.WriteByte('\n')

			for i := 0; i < depth; i++ {
				result.WriteByte('\t')
			}

		}

		result.WriteString(fmt.Sprintf("</%s>", node.Data))

	case html.TextNode:

		content := bytes.TrimSpace([]byte(node.Data))

		if len(content) > 0 {

			result.WriteByte('\n')

			for i := 0; i < depth; i++ {
				result.WriteByte('\t')
			}

			result.Write(content)

		}

	case html.CommentNode:

		result.WriteByte('\n')

		for i := 0; i < depth; i++ {
			result.WriteByte('\t')
		}

		result.WriteString(fmt.Sprintf("<!-- %s -->", node.Data))

	case html.RawNode:

		result.WriteByte('\n')

		for i := 0; i < depth; i++ {
			result.WriteByte('\t')
		}

		result.WriteString(node.Data)

	}

}
