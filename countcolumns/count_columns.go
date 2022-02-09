package countcolumns

import (
	"encoding/csv"
	"io"
)

func CountColumnsCSVReader(rd io.Reader, delimiter rune) (int, error) {
	csvReader := csv.NewReader(rd)
	csvReader.Comma = delimiter
	record, err := csvReader.Read()
	if err!=nil {
		return 0, err
	}
	return len(record), nil
}