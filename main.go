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

type Opts struct {
	canonical bool
	verbose   bool
}

type Line struct {
	offset int
	hex    []string
	ascii  string
}

func (l Line) String(opts Opts) string {
	var line string

	if opts.canonical {
		hex1 := strings.Join(l.hex[:8], " ")
		hex2 := strings.Join(l.hex[8:], " ")

		line = fmt.Sprintf("%08x  %-24s %-24s |%s|", l.offset, hex1, hex2, l.ascii)
	} else {
		var hex strings.Builder
		for chunk := range slices.Chunk(l.hex, 2) {
			part := strings.TrimSpace(chunk[1] + chunk[0])
			if part == "" {
				hex.WriteString("    ")
			} else if len(part) < 4 {
				fmt.Fprintf(&hex, "%04s", part)
			} else {
				hex.WriteString(part)
			}
			hex.WriteString(" ")
		}

		line = fmt.Sprintf("%07x %s", l.offset, strings.TrimSuffix(hex.String(), " "))
	}

	return line
}

type Dump struct {
	lines  []Line
	offset int
}

func (d Dump) Render(w io.Writer, opts Opts) {
	dupe := false

	for i, line := range d.lines {
		if i > 0 && !opts.verbose {
			if slices.Equal(line.hex, d.lines[i-1].hex) {
				dupe = true
				continue
			}
			if dupe {
				dupe = false
				fmt.Fprintln(w, "*")
			}
		}
		fmt.Fprintln(w, line.String(opts))
	}

	if opts.canonical {
		fmt.Fprintf(w, "%08x\n", d.offset)
	} else {
		fmt.Fprintf(w, "%07x\n", d.offset)
	}
}

func parse(input []byte) Dump {
	lines := []Line{}
	offset := 0

	for chunk := range slices.Chunk(input, 16) {
		line := Line{
			offset: offset,
			hex:    slices.Repeat([]string{"  "}, 16),
			ascii:  "",
		}

		for i := range chunk {
			line.hex[i] = fmt.Sprintf("%02x", chunk[i])

			value, _ := hex.DecodeString(string(fmt.Sprintf("%02x", chunk[i])))
			if value[0] < 32 || value[0] > 126 { // non-printable ASCII range
				line.ascii += "."
			} else {
				line.ascii += string(value)
			}
		}
		lines = append(lines, line)
		offset += len(chunk)
	}

	return Dump{lines, offset}
}

func main() {
	canonical := flag.Bool("C", false, "canonical hex+ASCII display ")
	verbose := flag.Bool("v", false, "display all input data")

	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage: gex [-C -v] <file>\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	inputFile := flag.Arg(0)
	if inputFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	opts := Opts{verbose: *verbose, canonical: *canonical}

	dump := parse(input)
	dump.Render(os.Stdout, opts)
}
