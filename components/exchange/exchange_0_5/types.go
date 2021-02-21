package exchange_0_5

type Field struct {
	Name   string
	Type   string
	Format string
	Tags   []string
}

type Structure []Field

type Data []interface{}
