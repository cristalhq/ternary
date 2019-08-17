# ternary
[![Build Status][travis-img]][travis-url]
[![GoDoc][doc-img]][doc-url]
[![Go Report Card][reportcard-img]][reportcard-url]

Ternary logic for Go.

## Install

```
go get github.com/cristalhq/ternary
```

## Example

```go
a, b := ternary.False, ternary.True
ternary.Not(a)
ternary.Imp(a, b)
ternary.MA(b)
```

## License

[MIT License](LICENSE).

[travis-img]: https://travis-ci.org/cristalhq/ternary.svg?branch=master
[travis-url]: https://travis-ci.org/cristalhq/ternary
[doc-img]: https://godoc.org/github.com/cristalhq/ternary?status.svg
[doc-url]: https://godoc.org/github.com/cristalhq/ternary
[reportcard-img]: https://goreportcard.com/badge/cristalhq/ternary
[reportcard-url]: https://goreportcard.com/report/cristalhq/ternary
