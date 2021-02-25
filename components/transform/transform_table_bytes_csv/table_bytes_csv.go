package transform_table_bytes_csv

import (
	"fmt"
	"strings"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"
	"github.com/pavlo67/data_exchange/components/extraction"

	"github.com/pavlo67/data_exchange/components/transform"
)

var _ transform.Operator = &transformStructureDataTable{}

type transformStructureDataTable struct {
	table     transform.Table
	path      string
	separator string
}

const onNew = "on transformStructureDataTable.New(): "

func New(path, separator string) (transform.Operator, error) {

	if separator == "" {
		return nil, fmt.Errorf("no 'separator' value")
	}
	transformOp := transformStructureDataTable{
		path:      strings.TrimSpace(path),
		separator: separator,
	}
	return &transformOp, nil
}

func (transformOp *transformStructureDataTable) Reset() error {
	transformOp.table = nil
	return nil
}

const onStat = "on transformStructureDataTable.Stat(): "

func (transformOp *transformStructureDataTable) Stat(params common.Map) error {
	return common.ErrNotImplemented
}

const onIn = "on transformStructureDataTable.In(): "

func (transformOp *transformStructureDataTable) In(selector *selectors.Term, data interface{}) error {
	var err error

	if data != nil {
		if bytes, _ := data.([]byte); bytes != nil {
			transformOp.table, err = extraction.TableBytes(bytes, transformOp.separator)
			if err != nil {
				return errors.CommonError(err, onIn)
			}
			return nil
		}

		return fmt.Errorf("wrong data to import: %#v", data)
	}

	if transformOp.path != "" {
		_, transformOp.table, err = extraction.TableFile(transformOp.path, transformOp.separator)
		if err != nil {
			return errors.CommonError(err, onIn)
		}
		return nil
	}

	return fmt.Errorf("no data and no file path to import")
}

func (transformOp *transformStructureDataTable) Out(selector *selectors.Term) (data interface{}, err error) {
	return transformOp.table, nil
}
