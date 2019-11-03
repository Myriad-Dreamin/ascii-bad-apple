package main

import (
	"flag"
)

var (
	readFlag = flag.Bool("read", false, "read from file")
)

func main() {
	if *readFlag {
		Read()
	} else {
		Write()
	}
}
