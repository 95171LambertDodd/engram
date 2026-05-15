package main

import (
	"fmt"
	"os"

	"github.com/engram/engram/cmd"
)

// version is set at build time via ldflags
var version = "dev"

func main() {
	if err := cmd.Execute(version); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
