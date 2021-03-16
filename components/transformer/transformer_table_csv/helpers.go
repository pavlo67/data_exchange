package transformer_table_csv

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pavlo67/data_exchange/components/structures"
)

func TableFile(filename, separator string) ([]byte, *structures.Table, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("reading %s got %s", filename, err)
	}

	table, err := TableString(string(data), separator)
	return data, table, nil
}

func TableString(data string, separator string) (*structures.Table, error) {
	if separator == "" {
		return nil, errors.New("on extraction.TableString(): no fields separator")
	}

	lines := strings.Split(data, "\n")

	var table structures.Table
	for _, line := range lines {
		if line == "" {
			continue
		}
		table.Rows = append(table.Rows, strings.Split(line, separator))
	}

	return &table, nil
}
