package main

import "pwettify/formatters"
import "fmt"
import "os"
import "sort"
import "strings"

func showHelp() {

	extensions := make([]string, 0)

	for ext, _ := range formatters.Formatters {
		extensions = append(extensions, ext)
	}

	sort.Strings(extensions)

	fmt.Println("Usage:                pwettify /path/to/file.ext")
	fmt.Println("Supported Extensions: " + strings.Join(extensions, ", "))

	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("")
	fmt.Println("pwettify /tmp/raw.json > /tmp/prettified.json;")
	fmt.Println("pwettify /tmp/raw.html > /tmp/prettified.html;")
	fmt.Println("")

}

func main() {

	var file string
	var buffer []byte

	if len(os.Args) == 2 {

		tmp_file := strings.TrimSpace(os.Args[1])
		tmp_buffer, err := os.ReadFile(tmp_file)

		if err == nil {
			file = tmp_file
			buffer = tmp_buffer
		}

	}

	if file != "" && strings.Contains(file, ".") {

		tmp := strings.Split(file, ".")
		ext := tmp[len(tmp)-1]

		formatter, ok := formatters.Formatters[ext]

		if ok == true {

			if formatter.Supports(buffer) {

				formatted := formatter.Format(buffer)
				fmt.Print(string(formatted))

			}

		} else {

			var formatter formatters.Formatter
			var found bool

			for _, other := range formatters.Formatters {

				if other.Supports(buffer) {
					formatter = other
					found = true
					break
				}

			}

			if found == true {

				formatted := formatter.Format(buffer)
				fmt.Print(string(formatted))

			} else {
				fmt.Fprintln(os.Stderr, "Unsupported file extension: \"" + ext + "\"")
				os.Exit(1)
			}

		}

	} else {
		showHelp()
		os.Exit(1)
	}

}
