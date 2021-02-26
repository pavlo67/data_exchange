package transform

import (
	"time"

	"github.com/pavlo67/data_exchange/components/ns"
	"github.com/pavlo67/data_exchange/components/vcs"
)

// general table -------------------------------------------------------

type Table [][]string

// records -------------------------------------------------------------

type Content01 struct {
	Title    string      `json:",omitempty" bson:",omitempty"`
	Summary  string      `json:",omitempty" bson:",omitempty"`
	DataType string      `json:",omitempty" bson:",omitempty"`
	Data     string      `json:",omitempty" bson:",omitempty"`
	Embedded []Content01 `json:",omitempty" bson:",omitempty"`
	Tags     []string    `json:",omitempty" bson:",omitempty"`
}

type Record01 struct {
	IssuedID  ns.ID       // TODO: ba careful, IssuedID can't be empty
	OwnerID   ns.ID       `json:",omitempty" bson:",omitempty"`
	ViewerID  ns.ID       `json:",omitempty" bson:",omitempty"`
	Content   Content01   `json:",inline"    bson:",inline"`
	History   vcs.History `json:",omitempty" bson:",omitempty"`
	CreatedAt time.Time   `json:",omitempty" bson:",omitempty"`
	UpdatedAt *time.Time  `json:",omitempty" bson:",omitempty"`
}

type RecordsPack01 struct {
	Title     string
	Items     []Record01
	History   vcs.History
	CreatedAt time.Time
}

//func (ris *RecordsPack01) In(data []byte, path string) (filenames []string, err error) {
//	return nil, json.Unmarshal(data, ris)
//}
//
//func (ris RecordsPack01) Out(path string) (data []byte, filenames []string, err error) {
//	jsonBytes, err := json.Marshal(ris)
//	return jsonBytes, nil, err
//}
