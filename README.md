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

[build-img]: https://github.com/cristalhq/atomix/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/atomix/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/atomix
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/atomix
[reportcard-img]: https://goreportcard.com/badge/cristalhq/atomix
[reportcard-url]: https://goreportcard.com/report/cristalhq/atomix
[coverage-img]: https://codecov.io/gh/cristalhq/atomix/branch/master/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/atomix
