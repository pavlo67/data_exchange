package extractor_helpers

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Tab(filename string) ([]byte, [][]string, error) {
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
