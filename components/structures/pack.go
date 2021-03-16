package structures

import (
	"reflect"

	"github.com/pavlo67/common/common"
)

type PackDescription struct {
	ItemDescription `json:",inline"    bson:",inline"`
	Fields          `json:",omitempty" bson:",omitempty"`
	ErrorsMap       `json:",omitempty" bson:",omitempty"`
}

type Pack struct {
	PackDescription `            json:",inline"    bson:",inline"`
	Items           interface{} `json:",omitempty" bson:",omitempty"`
}

func (pack *Pack) Stat() PackStat {
	var packStat PackStat

	if pack == nil || pack.Items == nil {
		return packStat
	}

	packStat.ItemsStat.Errored = len(pack.ErrorsMap) // TODO??? check non empty pack.ErrorsMap values only
	if reflect.TypeOf(pack.Items).Kind() == reflect.Slice {
		v := reflect.ValueOf(pack.Items)
		packStat.ItemsStat.Total = v.Len()
		for i := 0; i < packStat.ItemsStat.Total; i++ {
			//itemI := v.Index(i).Interface()
			if !common.IsNil(v.Index(i).Interface()) {
				packStat.ItemsStat.NonEmpty++
			}
		}
	}

	packStat.FieldsStat = pack.Fields.Stat()
	packStat.ErrorsStat = pack.ErrorsMap.Stat()

	return packStat
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
