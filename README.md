# gofuckyourself

[![GoDoc](https://godoc.org/github.com/JoshuaDoes/gofuckyourself?status.svg)](https://godoc.org/github.com/JoshuaDoes/gofuckyourself)
[![Go Report Card](https://goreportcard.com/badge/github.com/JoshuaDoes/gofuckyourself)](https://goreportcard.com/report/github.com/JoshuaDoes/gofuckyourself)

A sanitization-based swear filter for Go.

# Installing
`go get github.com/JoshuaDoes/gofuckyourself`

# Example
```Go
package main

import (
	"fmt"

	swearfilter "github.com/JoshuaDoes/gofuckyourself"
)

var message = "This is a fucking message with shitty swear words."
var swears = []string{"fuck", "shit"}

func main() {
	filter := swearfilter.New(false, false, false, false, swears...)
	swearFound, swearsFound, err := filter.Check(message)
	fmt.Println("Swear found: ", swearFound)
	fmt.Println("Swears found: ", swearsFound)
	fmt.Println("Error: ", err)
}
```
### Output
```
> go run main.go
Swear found:  true
Swears found:  [fuck shit]
Error:  <nil>
```

## License
The source code for gofuckyourself is released under the MIT License. See LICENSE for more details.

## Donations
All donations are appreciated and help me stay awake at night to work on this more. Even if it's not much, it helps a lot in the long run!

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://paypal.me/JoshuaDoes)