# :art: brickcolor

[![Build Status](https://travis-ci.org/peterhellberg/brickcolor.svg?branch=master)](https://travis-ci.org/peterhellberg/brickcolor)
[![Go Report Card](https://goreportcard.com/badge/github.com/peterhellberg/brickcolor?style=flat)](https://goreportcard.com/report/github.com/peterhellberg/brickcolor)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/brickcolor)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/brickcolor#license-mit)

A (generated) Go package with all of the Roblox [BrickColor Codes](https://developer.roblox.com/articles/BrickColor-Codes)

## Installation

    go get -u github.com/peterhellberg/brickcolor

## Usage

```go
package main

import (
	"fmt"
	"math/rand"

	"github.com/peterhellberg/brickcolor"
)

func main() {
	rand.Seed(123)

	println(brickcolor.Random())
	println(brickcolor.Number(1032))
}

func println(bc brickcolor.BrickColor) {
	fmt.Println(bc.R, bc.G, bc.B, bc.Name)
}
```

This will output:

```
104 116 172 Medium bluish violet
255 0 191 Hot pink
```

## Generating a new version

```bash
$ go generate
```

## License (MIT)

Copyright (c) 2018-2019 [Peter Hellberg](https://c7.se)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
