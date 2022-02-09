package countlines

import "testing"

func TestCountLines(t *testing.T) {
	file := "count_lines.go"
	exp := 51

	lines, err := Count(file)
	if err != nil {
		t.Error(err)
	}
	if lines != exp {
		t.Errorf("file %s should have %d line, not %d lines", file, exp, lines)
	}
}
