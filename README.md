# ternary

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]

Ternary logic for Go.

## Install

```
go get github.com/cristalhq/ternary
```

## Example

```go
a, b, c := ternary.False, ternary.Unknown, ternary.True
ternary.Not(a)
ternary.Imp(a, b)
ternary.ImpL(a, c)
ternary.MA(b)
```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/ternary/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/ternary/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/ternary
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/ternary
[reportcard-img]: https://goreportcard.com/badge/cristalhq/ternary
[reportcard-url]: https://goreportcard.com/report/cristalhq/ternary
[coverage-img]: https://codecov.io/gh/cristalhq/ternary/branch/master/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/ternary
