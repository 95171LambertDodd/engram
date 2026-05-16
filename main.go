package main

import (
	"fmt"
	"os"

	"github.com/engram/engram/cmd"
)

// version is set at build time via ldflags
// e.g. go build -ldflags "-X main.version=1.0.0"
var version = "dev"

func main() {
	if err := cmd.Execute(version); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
