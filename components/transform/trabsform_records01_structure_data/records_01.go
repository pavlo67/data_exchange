package transform_records_01

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/filelib"
	"github.com/pavlo67/common/common/selectors"
	"github.com/pavlo67/data_exchange/components/transform"
)

var _ transform.Operator = &transform01Files{}

type transform01Files struct {
	path                  string
	recordsExchangePack01 transform.RecordsExchangePack01
}

const onNew = "on transform01Files.New(): "

func New(path string) (transform.Operator, error) {
	correctedPath, err := filelib.Dir(path)
	if err != nil {
		return nil, errors.CommonError(err, onNew)
	}

	transformOp := transform01Files{
		path: correctedPath,
	}

	return &transformOp, nil
}

//func (transformOp *transform01Files) Name() string {
//	return string(InterfaceKey)
//}
//
//func (transformOp *transform01Files) Version() transform.Version {
//	return transform.Version("0.1.0")
//}

func (transformOp *transform01Files) Reset() error {
	transformOp.recordsExchangePack01 = transform.RecordsExchangePack01{}
	return nil
}

const onStat = "on transform01Files.Stat(): "

// from internal database
func (transformOp *transform01Files) Stat(params common.Map) error {
	return common.ErrNotImplemented
}

const onRead = "on transform01Files.Read(): "

// from internal database
func (transformOp *transform01Files) Read(selector *selectors.Term) error {

	var filename string
	// TODO read filename from selector

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf(onRead+": reading %s got %s", filename, err)
	}

	var recordsExchangePack transform.RecordsExchangePack01
	if err = json.Unmarshal(data, &recordsExchangePack); err != nil {
		return fmt.Errorf(onRead+": reading %s got %s", filename, err)
	}
	transformOp.recordsExchangePack01 = recordsExchangePack

	return nil
}

const onSave = "on transform01Files.Save()"

// into internal database
func (transformOp *transform01Files) Save(selector *selectors.Term) error {

	data, err := json.Marshal(transformOp.recordsExchangePack01)
	if err != nil {
		return fmt.Errorf(onSave+": marshalling data got %s", err)
	}

	var filename string
	// TODO read filename from selector

	if err = ioutil.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf(onSave+": writing into %s got %s", filename, err)
	}

	return nil
}

const onImport = "on transform01Files.In()"

// from external source
func (transformOp *transform01Files) In(selector *selectors.Term, data interface{}) error {
	switch v := data.(type) {
	case transform.RecordsExchangePack01:
		transformOp.recordsExchangePack01 = v
		return nil
	case *transform.RecordsExchangePack01:
		if v != nil {
			transformOp.recordsExchangePack01 = *v
			return nil
		}

		return errors.New(onImport + ": nil data for .In()")
	}

	return fmt.Errorf(onImport+": wrong data type (%T) for .In()", data)
}

// to external source
func (transformOp *transform01Files) Out(selector *selectors.Term) (data interface{}, err error) {
	return transformOp.recordsExchangePack01, nil
}
