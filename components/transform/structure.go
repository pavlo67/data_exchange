package transform

type Field struct {
	Name   string
	Type   string
	Format string
	Tags   []string
}

type Structure struct {
	Fields []Field
	Table  Table
}
