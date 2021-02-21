package exchange_0_1

import (
	"encoding/json"
	"time"

	"github.com/pavlo67/data_exchange/components/ns"
	"github.com/pavlo67/data_exchange/components/vcs"
)

type Content struct {
	Title    string    `json:",omitempty" bson:",omitempty"`
	Summary  string    `json:",omitempty" bson:",omitempty"`
	DataType string    `json:",omitempty" bson:",omitempty"`
	Data     string    `json:",omitempty" bson:",omitempty"`
	Embedded []Content `json:",omitempty" bson:",omitempty"`
	Tags     []string  `json:",omitempty" bson:",omitempty"`
}

type RecordItem struct {
	IssuedID  ns.ID       // TODO: ba careful, IssuedID can't be empty
	OwnerID   ns.ID       `json:",omitempty" bson:",omitempty"`
	ViewerID  ns.ID       `json:",omitempty" bson:",omitempty"`
	Content   Content     `json:",inline"    bson:",inline"`
	History   vcs.History `json:",omitempty" bson:",omitempty"`
	CreatedAt time.Time   `json:",omitempty" bson:",omitempty"`
	UpdatedAt *time.Time  `json:",omitempty" bson:",omitempty"`
}

type RecordsExchangePack struct {
	Title     string
	Items     []RecordItem
	History   vcs.History
	CreatedAt time.Time
}

func (ris *RecordsExchangePack) Import(data []byte, path string) (filenames []string, err error) {
	return nil, json.Unmarshal(data, ris)
}

func (ris RecordsExchangePack) Export(path string) (data []byte, filenames []string, err error) {
	jsonBytes, err := json.Marshal(ris)
	return jsonBytes, nil, err
}
