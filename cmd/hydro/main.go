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
	fmt.Println(buffer.ToString())

	buffer.InsertChar('a')
	buffer.InsertChar('1')
	buffer.InsertChar('1')
	buffer.InsertString(" hello world!")
	buffer.InsertString(" hello world!")
	buffer.InsertString(" hello world!")
	buffer.InsertString(" hello world!")
	buffer.InsertString(" hello world!")
	buffer.InsertString(" hello world!")
	buffer.InsertString(" hello world!")
	buffer.InsertString(" hello world!")
	buffer.InsertString(" hello world!")

	fmt.Println(buffer.ToString())
}
