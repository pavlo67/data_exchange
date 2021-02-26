package transform_records01_structure

import (
	"fmt"
	"time"

	"github.com/pavlo67/data_exchange/components/ns"

	"github.com/pavlo67/common/common/errors"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data_exchange/components/transform"
)

var _ transform.Operator = &transformRecords01Structure{}

type transformRecords01Structure struct {
	recordsPack01 transform.RecordsPack01
}

const onNew = "on transformRecords01Structure.New(): "

func New() (transform.Operator, error) {
	transformOp := transformRecords01Structure{}

	return &transformOp, nil
}

func (transformOp *transformRecords01Structure) Reset() error {
	transformOp.recordsPack01 = transform.RecordsPack01{}
	return nil
}

const onStat = "on transformRecords01Structure.Stat(): "

// from internal database
func (transformOp *transformRecords01Structure) Stat(params common.Map) error {
	return common.ErrNotImplemented
}

const onIn = "on transformRecords01Structure.In()"

// from external source
func (transformOp *transformRecords01Structure) In(selector *selectors.Term, data interface{}) error {
	if err := transformOp.Reset(); err != nil {
		return errors.CommonError(err, onIn)
	}

	var structure *transform.Structure

	switch v := data.(type) {
	case transform.Structure:
		structure = &v
		return nil
	case *transform.Structure:
		structure = v
	}

	if structure == nil {
		return fmt.Errorf(onIn+": wrong data (%#v)", data)

	}

	transformOp.recordsPack01 = transform.RecordsPack01{
		Title:     structure.Title,
		Items:     make([]transform.Record01, len(structure.Table)),
		History:   structure.History, // TODO!!! modify .History here
		CreatedAt: time.Now(),
	}

	for i, row := range structure.Table {
		var issuedID, ownerID, vieverID ns.ID

		transformOp.recordsPack01.Items[i] = transform.Record01{
			IssuedID: "",
			OwnerID:  "",
			ViewerID: "",
			Content: transform.Content01{
				Title:    "",
				Summary:  "",
				DataType: "",
				Data:     "",
				Embedded: nil,
				Tags:     nil,
			},
			History:   nil,
			CreatedAt: time.Time{},
			UpdatedAt: nil,
		}

	}

	return common.ErrNotImplemented
}

//type Content01 struct {
//	Title    string    `json:",omitempty" bson:",omitempty"`
//	Summary  string    `json:",omitempty" bson:",omitempty"`
//	DataType string    `json:",omitempty" bson:",omitempty"`
//	Data     string    `json:",omitempty" bson:",omitempty"`
//	Embedded []Content01 `json:",omitempty" bson:",omitempty"`
//	Tags     []string  `json:",omitempty" bson:",omitempty"`
//}
//
//type Record01 struct {
//	IssuedID  ns.ID       // TODO: ba careful, IssuedID can't be empty
//	OwnerID   ns.ID       `json:",omitempty" bson:",omitempty"`
//	ViewerID  ns.ID       `json:",omitempty" bson:",omitempty"`
//	Content01   Content01     `json:",inline"    bson:",inline"`
//	History   vcs.History `json:",omitempty" bson:",omitempty"`
//	CreatedAt time.Time   `json:",omitempty" bson:",omitempty"`
//	UpdatedAt *time.Time  `json:",omitempty" bson:",omitempty"`
//}

// to external source
func (transformOp *transformRecords01Structure) Out(selector *selectors.Term) (data interface{}, err error) {
	return transformOp.recordsPack01, nil
}

//const onRead = "on transformRecords01Structure.Read(): "
//
//// from internal database
//func (transformOp *transformRecords01Structure) Read(selector *selectors.Term) error {
//
//	var filename string
//	// TODO read filename from selector
//
//	data, err := ioutil.ReadFile(filename)
//	if err != nil {
//		return fmt.Errorf(onRead+": reading %s got %s", filename, err)
//	}
//
//	var recordsExchangePack transform.RecordsPack01
//	if err = json.Unmarshal(data, &recordsExchangePack); err != nil {
//		return fmt.Errorf(onRead+": reading %s got %s", filename, err)
//	}
//	transformOp.recordsPack01 = recordsExchangePack
//
//	return nil
//}
//
//const onSave = "on transformRecords01Structure.Save()"
//
//// into internal database
//func (transformOp *transformRecords01Structure) Save(selector *selectors.Term) error {
//
//	data, err := json.Marshal(transformOp.recordsPack01)
//	if err != nil {
//		return fmt.Errorf(onSave+": marshalling data got %s", err)
//	}
//
//	var filename string
//	// TODO read filename from selector
//
//	if err = ioutil.WriteFile(filename, data, 0644); err != nil {
//		return fmt.Errorf(onSave+": writing into %s got %s", filename, err)
//	}
//
//	return nil
//}
