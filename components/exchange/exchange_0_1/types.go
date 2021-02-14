package exchange_0_1

import (
	"encoding/json"

	"github.com/pavlo67/common/common"
)

type ID common.IDStr

type RecordItem struct {
	ID ID `json:"id"          bson:"_id"`
}

type RecordItems struct {
	Title string
	Items []RecordItem
}

func (ris *RecordItems) Import(data []byte, path string) (filenames []string, err error) {
	return nil, json.Unmarshal(data, ris)
}

func (ris RecordItems) Export(path string) (data []byte, filenames []string, err error) {
	jsonBytes, err := json.Marshal(ris)
	return jsonBytes, nil, err
}
