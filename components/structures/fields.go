package structures

import "fmt"

type Field struct {
	Name   string   `json:",omitempty" bson:",omitempty"`
	Type   string   `json:",omitempty" bson:",omitempty"`
	Format string   `json:",omitempty" bson:",omitempty"`
	Tags   []string `json:",omitempty" bson:",omitempty"`
}

type Fields []Field

func (fields Fields) Index() (map[string]int, error) {
	index := map[string]int{}

	for i, f := range fields {
		if _, ok := index[f.Name]; ok {
			return nil, fmt.Errorf("duplicate field '%s' in fields %#v", f.Name, fields)
		}
		index[f.Name] = i
	}

	return index, nil
}
