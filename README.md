
# pwettify

Standalone formatter for the command line, that runs on the target
system without having to install an interpreter or any external
runtime dependencies.


# Motivation

I got annoyed by Perl's and Python's fucked up ecosystem for simple
tools like prettifiers, where an installation on a system took me
half an hour managing outdated dependencies and packages on AUR.

This project is my attempt of supporting a couple of data formats
that I need for reverse engineering purposes on a daily basis, the
idea is to keep the codebase small and use upstream decoders/encoders
wherever possible (including golang.org/x).

If there's a file format that is currently unsupported, don't hesitate
to create an issue! Pull requests are also welcomed here! :)


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
# Usage
pwettify /path/to/raw.json > /path/to/prettified.json;
```


# Supported Formats

Some formats are planned but unsupported right now. If you have a nice
way to support them, please don't hesitate to make a pull request!

| Extension | Formatter                 | Package                                                      |
|:---------:|:--------------------------|:-------------------------------------------------------------|
| csv       |                           |                                                              |
| ini       |                           |                                                              |
| json      | [json](./formatters/json) | `encoding/json`                                              |
| htm       | [html](./formatters/html) | [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net) |
| html      | [html](./formatters/html) | [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net) |
| xml       | [xml](./formatters/xml)   | `encoding/xml`                                               |
| xhtml     | [xml](./formatters/xml)   | `encoding/xml`                                               |
| yaml      | [yaml](./formatters/yaml) | [gopkg.in/yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3)      |


# License

AGPL3

