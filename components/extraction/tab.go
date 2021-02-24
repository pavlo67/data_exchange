package extraction

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pavlo67/data_exchange/components/exchange"
)

func Tab(filename string) ([]byte, exchange.TabbedData, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("reading %s got %s", filename, err)
	}

	lines := strings.Split(string(data), "\n")

	var tab [][]string
	for _, line := range lines {
		tab = append(tab, strings.Split(line, "\t"))
	}

	return data, tab, nil
}
