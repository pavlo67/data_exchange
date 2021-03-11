package structures

import (
	"fmt"
	"time"

	"github.com/pavlo67/data_exchange/components/vcs"
)

type Rows [][]string

type Table struct {
	Title     string      `json:",omitempty" bson:",omitempty"`
	Fields    Fields      `json:",omitempty" bson:",omitempty"`
	Data      Rows        `json:",omitempty" bson:",omitempty"`
	ErrorsMap ErrorsMap   `json:",omitempty" bson:",omitempty"`
	History   vcs.History `json:",omitempty" bson:",omitempty"`
	CreatedAt time.Time   `json:",omitempty" bson:",omitempty"`
	UpdatedAt *time.Time  `json:",omitempty" bson:",omitempty"`
}

func (table *Table) Stat() (*TableStat, error) {
	if table == nil {
		return nil, nil
	}

	var tableStat TableStat
	tableStat.RowsStat.Total = len(table.Data)
	tableStat.RowsStat.Errored = len(table.ErrorsMap) // TODO??? check non empty pack.ErrorsMap values only

	tableStat.FieldsStat = make(FieldsStat, len(table.Fields)+1)
	tableStat.ColumnsStat = FieldsStat{}

	for _, row := range table.Data {
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
