package formatters

import "pwettify/formatters/html"
import "pwettify/formatters/json"
import "pwettify/formatters/xml"
import "pwettify/formatters/yaml"

type Formatter struct {
	Supports func(buffer []byte) bool
	Format   func(buffer []byte) []byte
}

var Formatters map[string]Formatter = map[string]Formatter{
	"htm":   {html.Supports, html.Format},
	"html":  {html.Supports, html.Format},
	"json":  {json.Supports, json.Format},
	"xml":   {xml.Supports,  xml.Format},
	"xhtml": {xml.Supports,  xml.Format},
	"yaml":  {yaml.Supports, yaml.Format},
}
