package extraction

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pavlo67/common/common/errors"

	"github.com/pavlo67/data_exchange/components/transform"
)

func TableFile(filename, separator string) ([]byte, transform.Table, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("reading %s got %s", filename, err)
	}

	table, err := TableBytes(data, separator)
	return data, table, nil
}

func TableBytes(data []byte, separator string) (transform.Table, error) {
	if separator == "" {
		return nil, errors.New("on extraction.TableBytes(): no fields separator")
	}

	lines := strings.Split(string(data), "\n")

	var table transform.Table
	for _, line := range lines {
		table = append(table, strings.Split(line, separator))
	}

	return table, nil
}
