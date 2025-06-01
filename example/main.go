package main

import (
	//Optionally name the import to hide the fuck in your later invocations, or maintain a fork!
	swearfilter "github.com/JoshuaDoes/gofuckyourself"

	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("must provide quoted message to check!")
	}
	message := os.Args[1]
	if len(os.Args) < 3 {
		panic("must provide one or more swears to use (with spaces between each)!")
	}
	swears := os.Args[2:]

	filter := swearfilter.NewSwearFilter(true, swears...)
	tripped, err := filter.Check(message)
	if err != nil {
		panic(err)
	}

	fmt.Println("Message:", message)
	fmt.Println("Swears:", tripped)
}
