package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

const LINE_LENGTH = 60

func main() {
	flag.Parse()

	inputFile := flag.Arg(0)
	if inputFile == "" {
		fmt.Println("Usage: gex <input file>")
		os.Exit(1)
	}

	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	displacement := 0
	for chunk := range slices.Chunk(input, 16) {
		var line strings.Builder
		fmt.Fprintf(&line, "%08x  ", displacement)

		var ascii strings.Builder
		for i := range chunk {
			if i == 8 {
				line.WriteString(" ")
			}
			fmt.Fprintf(&line, "%02x", chunk[i])

			value, _ := hex.DecodeString(string(fmt.Sprintf("%02x", chunk[i])))
			if value[0] < 32 || value[0] > 126 { // non-printable ASCII range
				ascii.WriteString(".")
			} else {
				ascii.WriteString(string(value))
			}

			line.WriteString(" ")
		}
		if line.Len() < LINE_LENGTH {
			line.WriteString(strings.Repeat(" ", LINE_LENGTH-line.Len()))
		}
		fmt.Fprintf(&line, "|%s|", ascii.String())
		fmt.Println(line.String())
		displacement += len(chunk)
	}

	if (displacement % 16) != 0 {
		fmt.Printf("%08x\n", displacement)
	}
}
