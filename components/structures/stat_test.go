package structures

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/common/common/errors"
)

func TestTableStat(t *testing.T) {
	table := Table{
		PackDescription: &PackDescription{
			ItemDescription: ItemDescription{
				Label:     "abc",
				CreatedAt: time.Now(),
			},
			Fields: []Field{
				{Name: "f1", Type: "t1", Format: "ff1", Tags: []string{"a", "b"}},
				{Name: "f2", Type: "t2", Format: "ff2", Tags: []string{"a", "b"}},
			},
			ErrorsMap: ErrorsMap{2: map[string]errors.Error{"f2": errors.CommonError("1", "2")}},
		},
		Rows: Rows{
			{"a", "b", "c", "d", "e"},
			{"a", "b", "c", "d", "e", "f"},
			{},
			{},
			{"1"},
		},
	}

	tableStat, err := table.Stat()
	require.NoError(t, err)
	require.NotNil(t, tableStat)

	t.Logf("%s", tableStat.String())

}

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
			{
				Name: "aaa",
				ItemsStat: ItemsStat{
					Total:    10,
					NonEmpty: 3,
					Errored:  3,
				},
			},
			{
				Name: "bbb",
				ItemsStat: ItemsStat{
					Total:    20,
					NonEmpty: 10,
					Errored:  10,
				},
			},

			{
				Name: "ccc",
				ItemsStat: ItemsStat{
					Total:    20,
					NonEmpty: 10,
					Errored:  10,
				},
			},
		},
		ColumnsStat: FieldsStat{
			{
				Name: "1",
				ItemsStat: ItemsStat{
					Total:    10,
					NonEmpty: 3,
					Errored:  3,
				},
			},

			{
				Name: "2",
				ItemsStat: ItemsStat{
					Total:    20,
					NonEmpty: 10,
					Errored:  10,
				},
			},

			{
				Name: "3",
				ItemsStat: ItemsStat{
					Total:    20,
					NonEmpty: 10,
					Errored:  10,
				},
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

func TestPackStat_String(t *testing.T) {
	packStat := PackStat{
		ItemsStat: ItemsStat{
			Total:    123,
			NonEmpty: 12,
			Errored:  333,
		},
		FieldsStat: FieldsStat{
			{
				Name: "aaa",
				ItemsStat: ItemsStat{
					Total:    10,
					NonEmpty: 3,
					Errored:  3,
				},
			},
			{
				Name: "bbb",
				ItemsStat: ItemsStat{
					Total:    20,
					NonEmpty: 10,
					Errored:  10,
				},
			},

			{
				Name: "ccc",
				ItemsStat: ItemsStat{
					Total:    20,
					NonEmpty: 10,
					Errored:  10,
				},
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

	t.Logf("%s", packStat.String())
}
