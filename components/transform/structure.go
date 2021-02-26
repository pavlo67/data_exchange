package transform

import "github.com/pavlo67/data_exchange/components/vcs"

type Field struct {
	Name   string
	Type   string
	Format string
	Tags   []string
}

type Structure struct {
	Title   string
	Fields  []Field
	Table   Table
	History vcs.History
}
