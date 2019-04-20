# ternary [![GoDoc][doc-img]][doc] [![Go Report Card][reportcard-img]][reportcard]

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

[doc-img]: https://godoc.org/github.com/cristalhq/ternary?status.svg
[doc]: https://godoc.org/github.com/cristalhq/ternary
[reportcard-img]: https://goreportcard.com/badge/cristalhq/ternary
[reportcard]: https://goreportcard.com/report/cristalhq/ternary
