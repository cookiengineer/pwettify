
# pwettify

Standalone prettifier / formatter for the command line, which runs
on the target machine without having to install additional packages
and/or dependencies.


# Motivation

I got annoyed by Perl's and Python's messed up ecosystem for simple
tools like prettifiers, where an installation on a system took me
more than half an hour of managing outdated dependencies and
packages on AUR.

This project is my attempt of supporting the most important file formats
in terms of de-minifying them which I need for reverse engineering apps
and programs on a daily basis.

The idea is to keep the codebase as small as possible and to use the
standard libraries wherever possible.

If there's a file format that is currently unsupported, don't hesitate
to create an issue! Pull requests are also very welcomed here! :)


## Building

```bash
cd /path/to/pwettify;

go build ./cmds/pwettify/main.go -o pwettify;
sudo cp pwettify /usr/bin/pwettify;
sudo chmod +x /usr/bin/pwettify;
```


## Usage

Running the `pwettify` command without a file parameter will show a
list of currently supported file extensions.

```bash
# Show Help
pwettify;

# Example Usage
pwettify /path/to/raw.json > /path/to/prettified.json;
```


# Supported Formats

Some file formats are planned but unsupported right now. If you have
a nice way to support them in pure go, please don't hesitate to make
a pull request!

| Extension | Formatter                 | Package                                                                         |
|:---------:|:--------------------------|:--------------------------------------------------------------------------------|
| css       | [css](./formatters/css)   | [github.com/evanw/esbuild](https://pkg.go.dev/github.com/evanw/esbuild/pkg/api) |
| csv       |                           |                                                                                 |
| ini       |                           |                                                                                 |
| js        | [js](./formatters/js)     | [github.com/evanw/esbuild](https://pkg.go.dev/github.com/evanw/esbuild/pkg/api) |
| json      | [json](./formatters/json) | `encoding/json`                                                                 |
| jsx       | [jsx](./formatters/jsx)   | [github.com/evanw/esbuild](https://pkg.go.dev/github.com/evanw/esbuild/pkg/api) |
| htm       | [html](./formatters/html) | [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html)               |
| html      | [html](./formatters/html) | [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html)               |
| manifest  | [json](./formatters/json) | `encoding/json`                                                                 |
| mjs       | [mjs](./formatters/mjs)   | [github.com/evanw/esbuild](https://pkg.go.dev/github.com/evanw/esbuild/pkg/api) |
| php       |                           |                                                                                 |
| sql       |                           |                                                                                 |
| ts        | [ts](./formatters/ts)     | [github.com/evanw/esbuild](https://pkg.go.dev/github.com/evanw/esbuild/pkg/api) |
| tsx       | [tsx](./formatters/tsx)   | [github.com/evanw/esbuild](https://pkg.go.dev/github.com/evanw/esbuild/pkg/api) |
| xml       | [xml](./formatters/xml)   | `encoding/xml`                                                                  |
| xhtml     | [xml](./formatters/xml)   | `encoding/xml`                                                                  |
| yaml      | [yaml](./formatters/yaml) | [gopkg.in/yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3)                         |


# License

AGPL3

