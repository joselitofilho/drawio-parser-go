<div align="center">

# drawio-parser-go

[![GitHub tag](https://img.shields.io/github/release/joselitofilho/drawio-parser-go?include_prereleases=&sort=semver&color=2ea44f&style=for-the-badge)](https://github.com/joselitofilho/drawio-parser-go/releases/)
[![Go Report Card](https://goreportcard.com/badge/github.com/joselitofilho/drawio-parser-go?style=for-the-badge)](https://goreportcard.com/report/github.com/joselitofilho/drawio-parser-go)
[![Code coverage](https://img.shields.io/badge/Coverage-100.0%25-2ea44f?style=for-the-badge)](#)

[![Made with Golang](https://img.shields.io/badge/Golang-1.21.6-blue?logo=go&logoColor=white&style=for-the-badge)](https://go.dev "Go to Golang homepage")
[![Using Diagrams](https://img.shields.io/badge/diagrams.net-orange?logo=&logoColor=white&style=for-the-badge)](https://app.diagrams.net/ "Go to Diagrams homepage")

[![BuyMeACoffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-ffdd00?style=for-the-badge&logo=buy-me-a-coffee&logoColor=black)](https://www.buymeacoffee.com/joselitofilho)

</div>

# Overview

This is a GoLang library designed to parse [draw.io][diagrams] diagram files and convert them into Go language 
structures. This project aims to provide a straightforward way to extract meaningful data from draw.io diagrams 
programmatically, enabling developers to integrate draw.io diagrams into their Go applications seamlessly.

## How to Use

```bash
$ go get github.com/joselitofilho/drawio-parser-go@latest
```

## Key Features

1. **Draw.io File Parsing**: Implement robust parsing functionality to extract data from draw.io XML files efficiently.
1. **Conversion to Go Structures**: Transform parsed data into Go language structures for easy manipulation and integration 
into Go projects.
1. **Support for Various Diagram Elements**: Support parsing various elements commonly found in draw.io diagrams, 
including shapes, connectors, text, and metadata.

## Potential Future Enhancements

1. **Export to Other Formats**: Extend functionality to support exporting parsed diagrams to various formats, such as 
image files or other diagram formats.

## Example Usage

`diagram.xml`:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<mxfile host="app.diagrams.net" modified="2024-02-26T17:40:23.051Z"
    agent="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
    etag="Ql3uf0GmtjkQuBo7mx8l" version="23.1.5" type="google">
    <diagram name="Page-1" id="oRquoGqMxknujh0Ahoql">
        <mxGraphModel dx="3372" dy="3488" grid="1" gridSize="10" guides="1" tooltips="1" connect="1"
            arrows="1" fold="1" page="1" pageScale="1" pageWidth="850" pageHeight="1100" math="0"
            shadow="0">
            <root>
                <mxCell id="kVijt7gfVD9ZtySMmpSK-1" value="myReceiver"
                    style="outlineConnect=0;dashed=0;verticalLabelPosition=bottom;verticalAlign=top;align=center;html=1;shape=mxgraph.aws3.lambda;fillColor=#F58534;gradientColor=none;"
                    parent="1" vertex="1">
                    <mxGeometry x="850" y="-1121.5" width="76.5" height="93" as="geometry" />
                </mxCell>
            </root>
        </mxGraphModel>
    </diagram>
</mxfile>
```

`main.go`:

```Go
package main

import (
	"fmt"

	"github.com/joselitofilho/drawio-parser-go/pkg/parser/xml"
)

func main() {
    diagram, err := xml.Parse("diagram.xml")
    fmt.Println(diagram, err)
}

```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to create an 
[issue][issues] or submit a pull request. Your contribution is much appreciated. See [Contributing](CONTRIBUTING.md).

[![open - Contributing](https://img.shields.io/badge/open-contributing-blue?style=for-the-badge)](CONTRIBUTING.md "Go to contributing")

## License

This project is licensed under the [MIT License](LICENSE).

[diagrams]: https://app.diagrams.net/
[issues]: https://github.com/diagram-code-generator/resources/issues