package structures

import "testing"

func TestTableStat_String(t *testing.T) {
	tableStat := TableStat{
		RowsStat: ItemsStat{
			Total:    123,
			NonEmpty: 12,
			Errored:  333,
		},
		//RowsValuesStat: ValuesStat{
		//	MinNonEmptyAmount: 33,
		//	MaxNonEmptyAmount: 44,
		//	MinNonEmptyIndex:  5,
		//	MaxNonEmptyIndex:  6,
		//},
		FieldsStat: FieldsStat{
			"aaa": ItemsStat{
				Total:    10,
				NonEmpty: 3,
				Errored:  3,
			},
			"bbb": ItemsStat{
				Total:    20,
				NonEmpty: 10,
				Errored:  10,
			},
			"ccc": ItemsStat{
				Total:    20,
				NonEmpty: 10,
				Errored:  10,
			},
		},
		ColumnsStat: FieldsStat{
			"1": ItemsStat{
				Total:    10,
				NonEmpty: 3,
				Errored:  3,
			},
			"2": ItemsStat{
				Total:    20,
				NonEmpty: 10,
				Errored:  10,
			},
			"3": ItemsStat{
				Total:    20,
				NonEmpty: 10,
				Errored:  10,
			},
		},
		ErrorsStat: ErrorsStat{
			Total:    1,
			Distinct: 1,
			Fields: map[string]int{
				"aaa": 3,
				"bbb": 10,
			},
		},
	}

	t.Logf("%s", tableStat.String())
}
