package structures

import (
	"time"

	"github.com/pavlo67/data_exchange/components/vcs"
)

type Rows [][]string

type Table struct {
	Title     string      `json:",omitempty" bson:",omitempty"`
	Fields    Fields      `json:",omitempty" bson:",omitempty"`
	Rows      Rows        `json:",omitempty" bson:",omitempty"`
	ErrorsMap ErrorsMap   `json:",omitempty" bson:",omitempty"`
	History   vcs.History `json:",omitempty" bson:",omitempty"`
	CreatedAt time.Time   `json:",omitempty" bson:",omitempty"`
	UpdatedAt *time.Time  `json:",omitempty" bson:",omitempty"`
}

func (table *Table) Stat() (*TableStat, error) {

	if table == nil {
		return nil, nil
	}

	//fieldsIndex, err := table.Fields.Index()
	//if err != nil {
	//	return nil, err
	//}

	var tableStat TableStat
	tableStat.RowsStat.Total = len(table.Rows)
	tableStat.RowsStat.Errored = len(table.ErrorsMap) // TODO??? check non empty pack.ErrorsMap values only

	tableStat.FieldsStat = map[string]ItemsStat{}
	tableStat.RowsValuesStat.MinNonEmptyIndex = -1

	for _, row := range table.Rows {
		if len(row) > 0 {
			tableStat.RowsStat.NonEmpty++
			nonEmptyAmount := 0
			nonEmptyIndexMin := -1
			nonEmptyIndexMax := -1
			for j, v := range row {
				var fieldName string
				if j < len(table.Fields) {
					fieldName = table.Fields[j].Name
				}

				fieldStat := tableStat.FieldsStat[fieldName]
				fieldStat.Total++

				if v != "" {
					fieldStat.NonEmpty++
					nonEmptyAmount++
					nonEmptyIndexMax = j
					if nonEmptyIndexMin < 0 {
						nonEmptyIndexMin = j
					}
				}

				tableStat.FieldsStat[fieldName] = fieldStat
			}

			if nonEmptyAmount > 0 {
				if tableStat.RowsValuesStat.MinNonEmptyAmount == 0 || nonEmptyAmount < tableStat.RowsValuesStat.MinNonEmptyAmount {
					tableStat.RowsValuesStat.MinNonEmptyAmount = nonEmptyAmount
				}
				if nonEmptyAmount > tableStat.RowsValuesStat.MaxNonEmptyAmount {
					tableStat.RowsValuesStat.MaxNonEmptyAmount = nonEmptyAmount
				}
				if tableStat.RowsValuesStat.MinNonEmptyIndex == -1 || nonEmptyIndexMin < tableStat.RowsValuesStat.MinNonEmptyIndex {
					tableStat.RowsValuesStat.MinNonEmptyIndex = nonEmptyIndexMin
				}
				if nonEmptyIndexMax > tableStat.RowsValuesStat.MaxNonEmptyIndex {
					tableStat.RowsValuesStat.MaxNonEmptyIndex = nonEmptyIndexMax
				}
			}
		}
	}

	tableStat.ErrorsStat = table.ErrorsMap.Stat()

	for _, f := range table.Fields {
		if tableStat.ErrorsStat.Fields[f.Name] > 0 {
			fieldStat := tableStat.FieldsStat[f.Name]
			fieldStat.Errored = tableStat.ErrorsStat.Fields[f.Name]
			tableStat.FieldsStat[f.Name] = fieldStat
		}
	}

	return &tableStat, nil
}
