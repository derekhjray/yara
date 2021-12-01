![Logo](/yara.png)

# yara

[![PkgGoDev](https://pkg.go.dev/badge/github.com/hillu/go-yara/v4)](https://pkg.go.dev/github.com/hillu/go-yara/v4)
[![Travis](https://travis-ci.org/hillu/go-yara.svg?branch=master)](https://travis-ci.org/hillu/go-yara)
[![Go Report Card](https://goreportcard.com/badge/github.com/hillu/go-yara)](https://goreportcard.com/report/github.com/hillu/go-yara)

This package is original forked from [go-yara](https://github.com/hillu/go-yara) repository.

Go bindings for [YARA](https://virustotal.github.io/yara/), staying as
close as sensible to the library's C-API while taking inspiration from
the `yara-python` implementation. And provides prebuild static libraries.

## Supported Prebuilt Libraries
- linux/amd64
- linux/arm64
- linux(musl)/amd64 (for alpine)
- linux(musl)/arm64 (for alpine)
- darwin/arm64
- `darwin/amd64 (incoming)`

## Build/Installation
### Dependencies
- libmagic
- libbz2
- liblzma
- libcrypto
- libjansson

### Generic Linux
Add v0.4.1 yara package to your go.mod file
```text
require (
	github.com/derekhjray/yara v0.4.1
)
```
And import the package where you need it
```go
import (
    "github.com/derekhjray/yara"
)
```
### Alpine Linux (musl)
Add v1.4.1 yara package to your go.mod file
```
require (
	github.com/derekhjray/yara v1.4.1
)
```
And import the package where you need it
```go
import (
    "github.com/derekhjray/yara"
)
```
### OSX Darwin
Either of ways above will be ok

### User Custom Libraries
The build tag `user_libs` can be used to tell the Go toolchain to use user customized
libraries. In this case, any compiler or linker
flags have to be set via the `CGO_CFLAGS` and `CGO_LDFLAGS`
environment variables, e.g.:
```
export CGO_CFLAGS="-I${YARA_SRC}/libyara/include"
export CGO_LDFLAGS="-L${YARA_SRC}/libyara/.libs -lyara"
go install -tags user_libs github.com/derekhjray/yara
```
…and rebuild your package.

## Version Definitions
`Major` represents prebuilt library type
- 0 presents `Generic Linux` and `Darwin` platform
- 1 presents `Alpine Linux` and `Darwin` platform

`Minor` represents _YAYA_'s Major version

`Patch` represents _YAYA_'s Minor version

_YARA_'s patch version will be omitted.

## License

BSD 2-clause, see LICENSE file in the source distribution.

## Author
Derek Ray <<derekhjray@gmail.com>>

Hilko Bengen <<bengen@hilluzination.de>>
