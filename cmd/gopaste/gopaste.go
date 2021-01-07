package main

import (
	"fmt"

	"github.com/kckrinke/go-clipboard"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Print(text)
}
