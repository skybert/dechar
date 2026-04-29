package main

import (
	"fmt"
	"os"
	"unicode/utf8"

	"golang.org/x/text/unicode/runenames"
)

func main() {
	for _, s := range os.Args[1:] {
		showCharDetails(s)
	}
}

func showCharDetails(s string) {
	for i, r := range s {
		fmt.Printf("Byte %d has code point %v U+%X %q\n", i, r, r, r)
		fmt.Printf("Code point %v is encoded using %v bytes\n", r, utf8.RuneLen(r))
		fmt.Printf("Code point %v is called %v\n", r, runenames.Name(r))
	}
}
