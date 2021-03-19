package structures

import (
	"fmt"
	"reflect"

	"github.com/pavlo67/common/common"
)

func Stat(pack Pack) PackStat {
	var packStat PackStat

	if pack == nil {
		return packStat
	}

	items := pack.Data()
	description := pack.Description()

	if reflect.TypeOf(items).Kind() == reflect.Slice {
		v := reflect.ValueOf(items)
		packStat.ItemsStat.Total = v.Len()
		for i := 0; i < packStat.ItemsStat.Total; i++ {
			//itemI := v.Index(i).Interface()
			if !common.IsNil(v.Index(i).Interface()) {
				packStat.ItemsStat.NonEmpty++
			}
		}
	}

	packStat.ItemsStat.Errored = len(description.ErrorsMap) // TODO??? check non empty pack.ErrorsMap values only
	packStat.FieldsStat = description.Fields.Stat()
	packStat.ErrorsStat = description.ErrorsMap.Stat()

	return packStat
}

var _ fmt.Stringer = &PackStat{}

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
		"\n  ItemsStat:\n                %s\n  FieldsStat: %s\n  ErrorsStat:\n                %s",
		packStat.ItemsStat.String(),
		packStat.FieldsStat.String(),
		packStat.ErrorsStat.String(),
	)
}

// !!! https://gist.github.com/pmn/5374494
// // Convert an interface{} containing a slice of structs into [][]string.
// func recordize(input interface{}) [][]string {
//	var records [][]string
//	var header []string // The first record in records will contain the names of the fields
//	object := reflect.ValueOf(input)
//
//	// The first record in the records slice should contain headers / field names
//	if object.Len() > 0 {
//		first := object.Index(0)
//		typ := first.Type()
//
//		for i := 0; i < first.NumField(); i++ {
//			header = append(header, typ.Field(i).Name)
//		}
//		records = append(records, header)
//	}
//
//	// Make a slice of objects to iterate through & populate the string slice
//	var items []interface{}
//	for i := 0; i < object.Len(); i++ {
//		items = append(items, object.Index(i).Interface())
//	}
//
//	// Populate the rest of the items into <records>
//	for _, v := range items {
//		item := reflect.ValueOf(v)
//		var record []string
//		for i := 0; i < item.NumField(); i++ {
//			itm := item.Field(i).Interface()
//			record = append(record, fmt.Sprintf("%v", itm))
//		}
//		records = append(records, record)
//	}
//	return records
// }
