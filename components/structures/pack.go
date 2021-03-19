package structures

type PackDescription struct {
	ItemDescription `json:",inline"    bson:",inline"`
	Fields          `json:",omitempty" bson:",omitempty"`
	ErrorsMap       `json:",omitempty" bson:",omitempty"`
}

type Pack interface {
	Description() PackDescription
	Data() interface{}
}
