# go-glob

[![Build Status](https://travis-ci.org/mgenware/go-glob.svg?branch=master)](http://travis-ci.org/mgenware/go-glob)

Another Glob Library for Golang

## Installation
```sh
go get github.com/mgenware/go-glob
```

## API
[![GoDoc](https://godoc.org/github.com/mgenware/go-glob?status.svg)](http://godoc.org/github.com/mgenware/go-glob)


## Example
```go
package main

import "github.com/mgenware/go-glob"

func main() {
	glob.IsFullMatch("abc", "a*")      // true
	glob.IsFullMatch("abc", "a?c")     // true
	glob.IsFullMatch("azzzzb", "a**b") // true
	glob.IsFullMatch("c?", "acd")      // false

	glob.IsPartialMatch("abc", "bc")     // true
	glob.IsPartialMatch("azzzzzb", "b*") // true
}
```
