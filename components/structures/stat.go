package structures

import (
	"fmt"
	"strings"
)

type ItemsStat struct {
	Total    int
	NonEmpty int
	Errored  int
}

func (itemsStat *ItemsStat) String() string {
	if itemsStat == nil {
		return "nil"
	}
	return fmt.Sprintf("Total:%4d, NonEmpty:%4d, Errored:%4d", itemsStat.Total, itemsStat.NonEmpty, itemsStat.Errored)

	//bytes, _ := json.Marshal(itemsStat)
	//return string(bytes)
}

type FieldStat struct {
	Name string
	ItemsStat
}

type FieldsStat []FieldStat

func (fieldsStat *FieldsStat) String() string {
	if fieldsStat == nil {
		return "nil"
	}

	var fieldsStatStr []string

	for _, f := range *fieldsStat {
		fieldsStatStr = append(fieldsStatStr, fmt.Sprintf("%-10s: %s", `"`+f.Name+`"`, f.ItemsStat.String()))
	}

	return "\n    " + strings.Join(fieldsStatStr, "\n    ")
}

type ErrorsStat struct {
	Total    int
	Distinct int
	Fields   map[string]int
}

func (errorsStat *ErrorsStat) String() string {
	if errorsStat == nil {
		return "nil"
	}
	//bytes, _ := json.Marshal(errorsStat)
	//return string(bytes)

	return fmt.Sprintf("Total:%4d, Distinct:%4d, Fields: %v", errorsStat.Total, errorsStat.Distinct, errorsStat.Fields)

}

type TableStat struct {
	RowsStat ItemsStat
	// RowsValuesStat ValuesStat
	FieldsStat
	ColumnsStat FieldsStat
	ErrorsStat
}

func (tableStat *TableStat) String() string {
	if tableStat == nil {
		return "nil"
	}
	return fmt.Sprintf(
		"\n  RowsStat:\n                %s\n  FieldsStat:%s\n  ColumnsStat:%s\n  ErrorsStat:\n                %s",
		tableStat.RowsStat.String(),
		// tableStat.RowsValuesStat.String(),
		tableStat.FieldsStat.String(),
		tableStat.ColumnsStat.String(),
		tableStat.ErrorsStat.String(),
	)
}
