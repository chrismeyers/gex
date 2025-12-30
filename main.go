package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Line struct {
	displacement int
	hex          []string
	ascii        string
}

func (l Line) String() string {
	hex1 := strings.Join(l.hex[:8], " ")
	hex2 := strings.Join(l.hex[8:], " ")

	return fmt.Sprintf("%08x  %-24s %-24s |%s|", l.displacement, hex1, hex2, l.ascii)
}

type Dump struct {
	lines        []Line
	displacement int
}

func (d Dump) Render(w io.Writer) {
	for _, line := range d.lines {
		fmt.Fprintln(w, line.String())
	}
	if (d.displacement % 16) != 0 {
		fmt.Fprintf(w, "%08x\n", d.displacement)
	}
}

func parse(input []byte) Dump {
	lines := []Line{}
	displacement := 0

	for chunk := range slices.Chunk(input, 16) {
		var line Line
		line.displacement = displacement

		for i := range chunk {
			line.hex = append(line.hex, fmt.Sprintf("%02x", chunk[i]))

			value, _ := hex.DecodeString(string(fmt.Sprintf("%02x", chunk[i])))
			if value[0] < 32 || value[0] > 126 { // non-printable ASCII range
				line.ascii += "."
			} else {
				line.ascii += string(value)
			}
		}
		if len(line.hex) < 16 {
			for i := 0; i < 16-len(line.hex); i++ {
				line.hex = append(line.hex, "  ")
			}
		}
		lines = append(lines, line)
		displacement += len(chunk)
	}

	return Dump{lines, displacement}
}

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

	dump := parse(input)
	dump.Render(os.Stdout)
}
