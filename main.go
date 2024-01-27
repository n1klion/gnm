package main

import (
	"fmt"
	gnm "gnm/cmd"
	"os"
)

func main() {
	app := gnm.Initialize()

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error running gnm: %s\n", err)
		os.Exit(1)
	}
}
