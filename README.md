termdeco
========

termdeco is a small library for cross-platform console text decoration for Go
language.

It provides functions for text decoration, Formatter, that can be used with
wrapper functions of fmt packages like

```go
termdeco.Println(termdeco.Red(v).BgGreen().Bold())
```

On *nix system, it just can be use with fmt package like

```go
fmt.Println(termdeco.Red(v).BgGreen().Bold())
```

## Installation

Install and update this go package with `go get -u github.com/tatsushid/termdeco`

## Examples

Import this package and write

```go
termdeco.Println(termdeco.Red("decorated").BgGreen().Bold())
```

It prints "decorated" as red, bold text (on Windows, bold is translated into
text brighter) on green background.

For more detail, refer [godoc][godoc]

## License
termdeco is under MIT License. See the [LICENSE][license] file for details.

[godoc]: http://godoc.org/github.com/tatsushid/termdeco
[license]: https://github.com/tatsushid/termdeco/blob/master/LICENSE
