package transformer_table_csv

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pavlo67/data_exchange/components/structures"
)

func RowsFile(filename, separator string) ([]byte, structures.Rows, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("reading %s got %s", filename, err)
	}

	table, err := RowsString(string(data), separator)
	return data, table, nil
}

func RowsString(data string, separator string) (structures.Rows, error) {
	if separator == "" {
		return nil, errors.New("on extraction.RowsString(): no fields separator")
	}

	lines := strings.Split(data, "\n")

	var rows structures.Rows
	for _, line := range lines {
		if line == "" {
			continue
		}
		rows = append(rows, strings.Split(line, separator))
	}

	return rows, nil
}
