package structures

import (
	"fmt"
)

type Rows [][]string

type Table struct {
	ItemDescription `json:",inline"    bson:",inline"`
	Fields          `json:",omitempty" bson:",omitempty"`
	Rows            `json:",omitempty" bson:",omitempty"`
	ErrorsMap       `json:",omitempty" bson:",omitempty"`
}

func (table *Table) Stat() (*TableStat, error) {
	if table == nil {
		return nil, nil
	}

	var tableStat TableStat
	tableStat.RowsStat.Total = len(table.Rows)
	tableStat.RowsStat.Errored = len(table.ErrorsMap) // TODO??? check non empty pack.ErrorsMap values only

	tableStat.FieldsStat = make(FieldsStat, len(table.Fields)+1)
	tableStat.ColumnsStat = FieldsStat{}

	for _, row := range table.Rows {
		if len(row) > 0 {
			tableStat.RowsStat.NonEmpty++
			for j, v := range row {
				fieldIndex := j
				if j >= len(table.Fields) {
					fieldIndex = len(table.Fields)
				}
				for j >= len(tableStat.ColumnsStat) {
					tableStat.ColumnsStat = append(tableStat.ColumnsStat, FieldStat{})
				}

				tableStat.FieldsStat[fieldIndex].Total++
				tableStat.ColumnsStat[j].Total++

				if v != "" {
					tableStat.FieldsStat[fieldIndex].NonEmpty++
					tableStat.ColumnsStat[j].NonEmpty++
				}
			}
		}
	}

	for j, field := range table.Fields {
		tableStat.FieldsStat[j].Name = field.Name
	}
	for j := range tableStat.ColumnsStat {
		tableStat.ColumnsStat[j].Name = fmt.Sprintf("%02d", j)
	}

	tableStat.ErrorsStat = table.ErrorsMap.Stat()

	for j, f := range table.Fields {
		if tableStat.ErrorsStat.Fields[f.Name] > 0 {
			tableStat.FieldsStat[j].Errored = tableStat.ErrorsStat.Fields[f.Name]
		}
	}

	return &tableStat, nil
}
