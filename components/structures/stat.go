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

type FieldsStat map[string]ItemsStat

func (fieldsStat *FieldsStat) String() string {
	if fieldsStat == nil {
		return "nil"
	}

	var fieldsStatStr []string
	for f, s := range *fieldsStat {
		fieldsStatStr = append(fieldsStatStr, `"`+f+`": {`+s.String()+"}")
	}

	return "\n    " + strings.Join(fieldsStatStr, "\n    ")
}

//type ValuesStat struct {
//	MinNonEmptyAmount int
//	MaxNonEmptyAmount int
//	MinNonEmptyIndex  int
//	MaxNonEmptyIndex  int
//}
//
//func (valuesStat *ValuesStat) String() string {
//	if valuesStat == nil {
//		return "nil"
//	}
//	//bytes, _ := json.Marshal(valuesStat)
//	//return string(bytes)
//	return fmt.Sprintf("MinNonEmptyAmount:%d, MaxNonEmptyAmount:%d, MinNonEmptyIndex:%d, MaxNonEmptyIndex:%d",
//		valuesStat.MinNonEmptyAmount, valuesStat.MaxNonEmptyAmount, valuesStat.MinNonEmptyIndex, valuesStat.MaxNonEmptyIndex)
//}

type PackStat struct {
	ItemsStat
	FieldsStat
	ErrorsStat
}

func (packStat *PackStat) String() string {
	if packStat == nil {
		return "nil"
	}
	return fmt.Sprintf(
		"\n  ItemsStat:\n    %s\n  FieldsStat: %s\n  ErrorsStat:\n    %s",
		packStat.ItemsStat.String(),
		packStat.FieldsStat.String(),
		packStat.ErrorsStat.String(),
	)
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

	return fmt.Sprintf("Total:%d, Distinct:%d, Fields: %v", errorsStat.Total, errorsStat.Distinct, errorsStat.Fields)

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
		"\n  RowsStat:\n    %s\n  FieldsStat:%s\n  ColumnsStat:%s\n  ErrorsStat:\n    %s", // RowsValuesStat: %s
		tableStat.RowsStat.String(),
		// tableStat.RowsValuesStat.String(),
		tableStat.FieldsStat.String(),
		tableStat.ColumnsStat.String(),
		tableStat.ErrorsStat.String(),
	)
}
