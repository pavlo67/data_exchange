package structures

type ItemsStat struct {
	Total    int
	NonEmpty int
	Errored  int
}

type FieldsStat map[string]ItemsStat

type ValuesStat struct {
	MinNonEmptyAmount int
	MaxNonEmptyAmount int
	MinNonEmptyIndex  int
	MaxNonEmptyIndex  int
}

type PackStat struct {
	ItemsStat
	FieldsStat
	ErrorsStat
}

type TableStat struct {
	RowsStat       ItemsStat
	RowsValuesStat ValuesStat
	FieldsStat
	ErrorsStat
}
