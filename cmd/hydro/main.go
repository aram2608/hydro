package main

import (
	"fmt"
	"hydro/buffer"
	"hydro/fileio"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [FILE PATH]\n", os.Args[0])
		os.Exit(1)
	}

	buffer := buffer.New(fileio.ReadFile(os.Args[1]))
	fmt.Print(buffer.ToString())
}
