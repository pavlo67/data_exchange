package transfer_table_csv

import (
	"fmt"
	"strings"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data/components/structures"
	"github.com/pavlo67/data/components/transfer"
)

var _ transfer.Operator = &transferTableCSV{}

type transferTableCSV struct {
	table *structures.Table
}

const onNew = "on transferTableCSV.New(): "

func New() (transfer.Operator, error) {
	return &transferTableCSV{}, nil
}

func (transferOp *transferTableCSV) reset() error {
	transferOp.table = nil
	return nil
}

func (transferOp *transferTableCSV) Name() string {
	return string(InterfaceKey)
}

const onIn = "on transferTableCSV.In(): "

func (transferOp *transferTableCSV) In(pack structures.Pack, params common.Map) error {
	if pack == nil {
		return errors.New(onIn + "nil pack to import")
	}

	if err := transferOp.reset(); err != nil {
		return errors.CommonError(err, onIn)
	}

	var err error

	separator := params.StringDefault("separator", "")
	if separator == "" {
		return fmt.Errorf(onIn + ": no 'separator' value in params")
	}

	transferOp.table = &structures.Table{
		PackDescription: pack.Description(),
	}

	data := pack.Data().Value()

	if data != nil {
		var dataStr string

		switch v := data.(type) {
		case []byte:
			dataStr = string(v)
		case *[]byte:
			if v == nil {
				return fmt.Errorf("nil data to import: %#v", data)
			}
			dataStr = string(*v)
		case string:
			dataStr = v
		case *string:
			if v == nil {
				return fmt.Errorf("nil data to import: %#v", data)
			}
			dataStr = *v
		default:
			return fmt.Errorf("wrong data to import: %#v", data)
		}

		// TODO!!! unescape separators

		transferOp.table.Rows, err = RowsString(dataStr, separator)
		if err != nil {
			return errors.CommonError(err, onIn)
		} else if transferOp.table == nil {
			return errors.CommonError("no table", onIn)
		}
		//transferOp.table.History = vcs.History{{
		//	Actor:  ns.URN(InterfaceKey), // TODO??????????????????????????????????????????
		//	Key:    vcs.CreatedAction,
		//	DoneAt: time.Now(),
		//}}
		return nil
	}

	if path := params.StringDefault("path", ""); path != "" {
		// TODO!!! unescape separators
		_, transferOp.table.Rows, err = RowsFile(path, separator)
		if err != nil {
			return errors.CommonError(err, onIn)
		}
		transferOp.table.Label = path
		//transferOp.table.History = vcs.History{{
		//	Actor:  ns.URN(InterfaceKey), // TODO??????????????????????????????????????????
		//	Key:    vcs.CreatedAction,
		//	DoneAt: time.Now(),
		//}}
		// TODO!!! add file info

		return nil
	}

	return fmt.Errorf("no data and no file path to import")
}

const onOut = "on transferTableCSV.Out(): "

func (transferOp *transferTableCSV) Out(selector *selectors.Term, params common.Map) (data structures.Pack, err error) {
	separator := params.StringDefault("separator", "")
	if separator == "" {
		return nil, fmt.Errorf(onOut + ": no 'separator' value in params")
	}

	var rowsStr []string

	if transferOp.table != nil {
		for _, row := range transferOp.table.Rows {
			// TODO!!! escape separators

			rowsStr = append(rowsStr, strings.Join(row, separator))
		}
	}

	return &structures.PackAny{
		PackDescription: transferOp.table.PackDescription,
		PackData:        structures.NewDataAny(strings.Join(rowsStr, "\n")),
	}, nil
}

const onStat = "on transferTableCSV.Stat(): "

func (transferOp *transferTableCSV) Stat(selector *selectors.Term, params common.Map) (interface{}, error) {
	return transferOp.table.Stat()
}

func (transferOp *transferTableCSV) Copy(selector *selectors.Term, params common.Map) (interface{}, error) {
	return transferOp.table, nil
}
