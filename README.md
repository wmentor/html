# HTML

[![Coverage Status](https://coveralls.io/repos/github/wmentor/html/badge.svg?branch=master&v=1.0.3)](https://coveralls.io/github/wmentor/html?branch=master)
[![https://goreportcard.com/report/github.com/wmentor/html](https://goreportcard.com/badge/github.com/wmentor/html)](https://goreportcard.com/report/github.com/wmentor/html)
[![https://pkg.go.dev/github.com/wmentor/html](https://pkg.go.dev/badge/github.com/wmentor/html.svg)](https://pkg.go.dev/github.com/wmentor/html)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Simple HTML parser and data fetcher library written on Golang under MIT License.

## Require

* Golang (version >= 1.20)
* golang.org/x/net

## Install

```
go get github.com/wmentor/html
```

## Usage

### Fetch data from URL

```golang
package main

import (
  "fmt"
  "time"

  "github.com/wmentor/html"
)

func main() {

  src := "https://edition.cnn.com"

  parser := html.New()

  opts := &html.GetOpts{
    Agent:"Mozilla/5.0 (compatible; MSIE 10.0)",
    Timeout: time.Second*60,
  }

  parser.Get(src,opts)
  fmt.Println( string(parser.Text()) )

  parser.EachLink(func(link string) {
    fmt.Println("url=" + link)
  } )

  parser.EachImage(func(link string) {
    fmt.Println("img=" + link)
  } )

  parser.EachIframe(func(link string) {
    fmt.Println("iframe=" + link)
  } )
}
```

### Fetch data from file/stdin

```golang
package main

import (
  "fmt"
  "os"

  "github.com/wmentor/html"
)

func main() {

  parser := html.New()

  parser.Parse(os.Stdin) // io.Reader
  fmt.Println( string(parser.Text()) )

  parser.EachLink(func(link string) {
    fmt.Println("url=" + link)
  } )

  parser.EachImage(func(link string) {
    fmt.Println("img=" + link)
  } )

  parser.EachIframe(func(link string) {
    fmt.Println("iframe=" + link)
  } )
}
```
