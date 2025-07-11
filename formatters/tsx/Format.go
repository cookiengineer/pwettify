package js

import "github.com/evanw/esbuild/pkg/api"
import "strings"

func Format(buffer []byte) []byte {

	result := api.Transform(string(buffer), api.TransformOptions{
		Loader:            api.LoaderTSX,
		MinifyWhitespace:  false,
		MinifyIdentifiers: false,
		MinifySyntax:      false,
		Format:            api.FormatDefault,
	})

	if len(result.Errors) != 0 {
		// Do Nothing?
	}

	lines := strings.Split(string(result.Code), "\n")

	for l, line := range lines {
		lines[l] = replacePrefix(line, "  ", "\t")
	}

	return []byte(strings.Join(lines, "\n"))

}

