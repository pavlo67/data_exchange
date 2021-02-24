package exchange_file_tabbed

import (
	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data_exchange/components/exchange"
	"github.com/pavlo67/data_exchange/components/extraction"
)

var _ exchange.Operator = &exchangeFileTabbed{}

type exchangeFileTabbed struct {
	path       string
	tabbedData exchange.TabbedData
}

const onNew = "on exchangeFileTabbed.New(): "

func New(path string) (exchange.Operator, error) {
	extractorOp := exchangeFileTabbed{
		path: path,
	}
	return &extractorOp, nil
}

func (exchangeOp *exchangeFileTabbed) Reset() error {
	exchangeOp.tabbedData = nil
	return nil
}

const onStat = "on exchangeFileTabbed.Stat(): "

func (exchangeOp *exchangeFileTabbed) Stat(params common.Map) error {
	return common.ErrNotImplemented
}

const onRead = "on exchangeFileTabbed.Read(): "

func (exchangeOp *exchangeFileTabbed) Read(selector *selectors.Term) error {
	_, tab, err := extraction.Tab(exchangeOp.path)
	if err != nil {
		return errors.CommonError(err, onRead)
	}

	// l.Infof("%s [%d bytes] --> %d lines", exchangeOp.path, len(data), len(tab))

	exchangeOp.tabbedData = tab

	return nil
}

const onSave = "on exchangeFileTabbed.Save(): "

func (exchangeOp *exchangeFileTabbed) Save(selector *selectors.Term) error {
	return errors.CommonError(common.ErrNotImplemented, onSave)
}

const onImport = "on exchangeFileTabbed.Import(): "

func (exchangeOp *exchangeFileTabbed) Import(selector *selectors.Term, structure, data interface{}) error {
	return errors.CommonError(common.ErrNotImplemented, onImport)
}

func (exchangeOp *exchangeFileTabbed) Export(selector *selectors.Term) (structure, data interface{}, err error) {
	tabbedData := exchangeOp.tabbedData
	exchangeOp.tabbedData = nil

	return nil, tabbedData, nil
}
