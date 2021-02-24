package exchange_records_01

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/filelib"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data_exchange/components/exchange"
)

var _ exchange.Operator = &exchange01Files{}

type exchange01Files struct {
	path                  string
	recordsExchangePack01 exchange.RecordsExchangePack01
}

const onNew = "on exchange01Files.New(): "

func New(path string) (exchange.Operator, error) {
	correctedPath, err := filelib.Dir(path)
	if err != nil {
		return nil, errors.CommonError(err, onNew)
	}

	exchangeOp := exchange01Files{
		path: correctedPath,
	}

	return &exchangeOp, nil
}

//func (exchangeOp *exchange01Files) Name() string {
//	return string(InterfaceKey)
//}
//
//func (exchangeOp *exchange01Files) Version() exchange.Version {
//	return exchange.Version("0.1.0")
//}

func (exchangeOp *exchange01Files) Reset() error {
	exchangeOp.recordsExchangePack01 = exchange.RecordsExchangePack01{}
	return nil
}

const onStat = "on exchange01Files.Stat(): "

// from internal database
func (exchangeOp *exchange01Files) Stat(params common.Map) error {
	return common.ErrNotImplemented
}

const onRead = "on exchange01Files.Read(): "

// from internal database
func (exchangeOp *exchange01Files) Read(selector *selectors.Term) error {

	var filename string
	// TODO read filename from selector

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf(onRead+": reading %s got %s", filename, err)
	}

	var recordsExchangePack exchange.RecordsExchangePack01
	if err = json.Unmarshal(data, &recordsExchangePack); err != nil {
		return fmt.Errorf(onRead+": reading %s got %s", filename, err)
	}
	exchangeOp.recordsExchangePack01 = recordsExchangePack

	return nil
}

const onSave = "on exchange01Files.Save()"

// into internal database
func (exchangeOp *exchange01Files) Save(selector *selectors.Term) error {

	data, err := json.Marshal(exchangeOp.recordsExchangePack01)
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

const onImport = "on exchange01Files.Import()"

// from external source
func (exchangeOp *exchange01Files) Import(selector *selectors.Term, structure, data interface{}) error {
	switch v := data.(type) {
	case exchange.RecordsExchangePack01:
		exchangeOp.recordsExchangePack01 = v
		return nil
	case *exchange.RecordsExchangePack01:
		if v != nil {
			exchangeOp.recordsExchangePack01 = *v
			return nil
		}

		return errors.New(onImport + ": nil data for .Import()")
	}

	return fmt.Errorf(onImport+": wrong data type (%T) for .Import()", data)
}

// to external source
func (exchangeOp *exchange01Files) Export(selector *selectors.Term) (structure, data interface{}, err error) {
	return nil, exchangeOp.recordsExchangePack01, nil
}
