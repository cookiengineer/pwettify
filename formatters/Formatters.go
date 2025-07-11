package formatters

import "pwettify/formatters/css"
import "pwettify/formatters/html"
import "pwettify/formatters/js"
import "pwettify/formatters/json"
import "pwettify/formatters/jsx"
import "pwettify/formatters/mjs"
import "pwettify/formatters/ts"
import "pwettify/formatters/tsx"
import "pwettify/formatters/xml"
import "pwettify/formatters/yaml"

type Formatter struct {
	Supports func(buffer []byte) bool
	Format   func(buffer []byte) []byte
}

var Formatters map[string]Formatter = map[string]Formatter{
	"css":      {css.Supports,  css.Format},
	"htm":      {html.Supports, html.Format},
	"html":     {html.Supports, html.Format},
	"js":       {js.Supports,   js.Format},
	"json":     {json.Supports, json.Format},
	"jsx":      {jsx.Supports,  jsx.Format},
	"manifest": {json.Supports, json.Format},
	"mjs":      {mjs.Supports,  mjs.Format},
	"ts":       {ts.Supports,   ts.Format},
	"tsx":      {tsx.Supports,  tsx.Format},
	"xml":      {xml.Supports,  xml.Format},
	"xhtml":    {xml.Supports,  xml.Format},
	"yaml":     {yaml.Supports, yaml.Format},
}
