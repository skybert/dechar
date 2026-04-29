package main

import (
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/spf13/pflag"
	"golang.org/x/text/unicode/runenames"
)

// Version is set at build time via: -ldflags "-X main.Version=<version>"
var Version = "foo"

var verbose = false
var showHelp = false
var showVersion bool

func init() {
	pflag.BoolVarP(&showHelp, "help", "h", false, "Don't panic")
	pflag.BoolVarP(&showVersion, "version", "e", false, "Show version")
	pflag.BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}

func main() {
	pflag.Parse()
	if showVersion {
		fmt.Printf("dechar version: %v\n", Version)
		os.Exit(0)
	}

	for _, s := range os.Args[1:] {
		if s[0] == '-' {
			continue
		}
		showCharDetails(s)

	}
}

func showCharDetails(s string) {
	for i, r := range s {
		if verbose {
			fmt.Printf("Byte %d has code point %v U+%X %q\n", i, r, r, r)
			fmt.Printf("Code point %v is encoded using %v bytes\n", r, utf8.RuneLen(r))
			fmt.Printf("Code point %v is called %v\n", r, runenames.Name(r))
		} else {
			fmt.Printf(
				"Byte pos: %d Code point: %v Name: %s Bytes: %d \n",
				i,
				r,
				runenames.Name(r),
				utf8.RuneLen(r))

		}
	}
}
