
# gofuckyourself

[![GoDoc](https://godoc.org/github.com/JoshuaDoes/gofuckyourself?status.svg)](https://godoc.org/github.com/JoshuaDoes/gofuckyourself)
[![Go Report Card](https://goreportcard.com/badge/github.com/JoshuaDoes/gofuckyourself)](https://goreportcard.com/report/github.com/JoshuaDoes/gofuckyourself)
[![cover.run](https://cover.run/go/github.com/JoshuaDoes/gofuckyourself.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2FJoshuaDoes%2Fgofuckyourself)

A sanitization-based swear filter for Go.

# Installing
`go get github.com/JoshuaDoes/gofuckyourself`

# Example
An up to date example can be found in the `example` subdirectory.

### Output
```
> git clone https://github.com/JoshuaDoes/gofuckyourself
...
> cd gofuckyourself
> go run ./example "This is a fûçking A S S of a message with shitty swear words." fuck shit ass
Message: This is a fûçking A S S of a message with shitty swear words.
Swears: [ass fuck shit]
```

## License
The source code for gofuckyourself is released under the MIT License. See LICENSE for more details.

## Donations
All donations are appreciated and help me stay awake at night to work on this more. Even if it's not much, it helps a lot in the long run!

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://paypal.me/JoshuaDoes)
