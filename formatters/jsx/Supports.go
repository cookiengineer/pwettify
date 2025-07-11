package js

import "github.com/evanw/esbuild/pkg/api"

func Supports(buffer []byte) bool {

	var result bool

	tmp := api.Transform(string(buffer), api.TransformOptions{
		Loader:            api.LoaderJSX,
		MinifyWhitespace:  false,
		MinifyIdentifiers: false,
		MinifySyntax:      false,
		Format:            api.FormatDefault,
	})

	if len(tmp.Errors) == 0 {
		result = true
	}

	return result

}
