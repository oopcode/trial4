package main

import (
	"flag"
	"fmt"
	"os"
	"trial4/external"
)

var (
	sortFlagSet = flag.NewFlagSet("sort", 0)
	input       = sortFlagSet.String("input", "example.txt", "Input file path")
	buffSize    = sortFlagSet.Int("buffer_size", 100, "Number of lines in a buffer")

	genFlagSet = flag.NewFlagSet("gen", 0)
	output     = genFlagSet.String("output", "example.txt", "Output file path")
	numLines   = genFlagSet.Int("num_lines", 50, "Number of output lines")
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "sort":
		sortFlagSet.Parse(os.Args[2:])
		external.Sort(*input, *buffSize)
	case "gen":
		genFlagSet.Parse(os.Args[2:])
		external.GenerateFile(*output, *numLines, 10)
	}
}
