package main

import (
	"fmt"
	"io"
	"os"

	"github.com/wodadehencou/tools/countlines"
)

func main() {
	readers := make([]io.Reader, 0)
	
	if len(os.Args) < 2 {
		readers = append(readers, os.Stdin)
		// read stdin
	} else {
		for _, f := range os.Args[1:] {
			if f == "-" {
				readers = append(readers, os.Stdin)
			} else {
				fp, err := os.Open(f)
				if err!=nil {
					fmt.Fprintf(os.Stderr, "Can not open file %s: %s", f, err)
					os.Exit(1)
				}
				defer fp.Close()
				readers = append(readers, fp)
			}
		}
	}
	
	r := io.MultiReader(readers...)
	
	lines, err := countlines.CountReader(r)
	if err!=nil {
					fmt.Fprintf(os.Stderr, "Count lines fail %s", err)
					os.Exit(1)
	}
	
	fmt.Fprintf(os.Stdout, "%d", lines)
	os.Exit(0)
}