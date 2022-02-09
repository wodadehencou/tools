package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/wodadehencou/tools/countcolumns"
)

var sep = flag.String("s", ",", "delimiter of csv file, [space|s] for space, [tab|t] for tab")

func init() {
	flag.Parse()
}

func main() {

	var reader io.Reader

	if len(flag.Args()) == 0 {
		reader = os.Stdin
	} else {
		if f := flag.Arg(0); f == "-" {
			reader = os.Stdin
		} else {
			fp, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can not open file %s: %s", f, err)
				os.Exit(1)
			}
			defer fp.Close()
			reader = fp
		}
	}

	var delimiter rune
	switch (strings.ToLower(*sep)) {
	case "tab", "t": delimiter = '\t';
	case "space", "s": delimiter = ' ';
	default: delimiter, _ = utf8.DecodeRuneInString(*sep)
	}
	cols, err := countcolumns.CountColumnsCSVReader(reader, delimiter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Count columns fail: %s", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%d", cols)
	os.Exit(0)
}
