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
	return fmt.Sprintf("Total:%d, NonEmpty:%d, Errored:%d", itemsStat.Total, itemsStat.NonEmpty, itemsStat.Errored)

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
		fieldsStatStr = append(fieldsStatStr, `"`+f+`":{`+s.String()+"}")
	}

	return strings.Join(fieldsStatStr, ", ")
}

type ValuesStat struct {
	MinNonEmptyAmount int
	MaxNonEmptyAmount int
	MinNonEmptyIndex  int
	MaxNonEmptyIndex  int
}

func (valuesStat *ValuesStat) String() string {
	if valuesStat == nil {
		return "nil"
	}
	//bytes, _ := json.Marshal(valuesStat)
	//return string(bytes)
	return fmt.Sprintf("MinNonEmptyAmount:%d, MaxNonEmptyAmount:%d, MinNonEmptyIndex:%d, MaxNonEmptyIndex:%d",
		valuesStat.MinNonEmptyAmount, valuesStat.MaxNonEmptyAmount, valuesStat.MinNonEmptyIndex, valuesStat.MaxNonEmptyIndex)

}

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
		"\nItemsStat:  %s\nFieldsStat: %s\nErrorsStat: %s",
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

	return fmt.Sprintf("Total:%d, Distinct:%d, Fields:%v", errorsStat.Total, errorsStat.Distinct, errorsStat.Fields)

}

type TableStat struct {
	RowsStat       ItemsStat
	RowsValuesStat ValuesStat
	FieldsStat
	ErrorsStat
}

func (tableStat *TableStat) String() string {
	if tableStat == nil {
		return "nil"
	}
	return fmt.Sprintf(
		"\nRowsStat:       %s\nRowsValuesStat: %s\nFieldsStat:     %s\nErrorsStat:     %s",
		tableStat.RowsStat.String(),
		tableStat.RowsValuesStat.String(),
		tableStat.FieldsStat.String(),
		tableStat.ErrorsStat.String(),
	)
}
