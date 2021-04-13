package transformer_table_csv

import (
	"fmt"
	"strings"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data_exchange/components/structures"
	"github.com/pavlo67/data_exchange/components/transformer"
)

var _ transformer.Operator = &transformTableCSV{}

type transformTableCSV struct {
	table *structures.Table
}

const onNew = "on transformTableCSV.New(): "

func New() (transformer.Operator, error) {
	return &transformTableCSV{}, nil
}

func (transformOp *transformTableCSV) reset() error {
	transformOp.table = nil
	return nil
}

func (transformOp *transformTableCSV) Name() string {
	return string(InterfaceKey)
}

const onIn = "on transformTableCSV.In(): "

func (transformOp *transformTableCSV) In(pack structures.Pack, params common.Map) error {
	if pack == nil {
		return errors.New(onIn + "nil pack to import")
	}

	if err := transformOp.reset(); err != nil {
		return errors.CommonError(err, onIn)
	}

	var err error

	separator := params.StringDefault("separator", "")
	if separator == "" {
		return fmt.Errorf(onIn + ": no 'separator' value in params")
	}

	transformOp.table = &structures.Table{
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

		transformOp.table.Rows, err = RowsString(dataStr, separator)
		if err != nil {
			return errors.CommonError(err, onIn)
		} else if transformOp.table == nil {
			return errors.CommonError("no table", onIn)
		}
		//transformOp.table.History = vcs.History{{
		//	Actor:  ns.URN(InterfaceKey), // TODO??????????????????????????????????????????
		//	Key:    vcs.CreatedAction,
		//	DoneAt: time.Now(),
		//}}
		return nil
	}

	if path := params.StringDefault("path", ""); path != "" {
		// TODO!!! unescape separators
		_, transformOp.table.Rows, err = RowsFile(path, separator)
		if err != nil {
			return errors.CommonError(err, onIn)
		}
		transformOp.table.Label = path
		//transformOp.table.History = vcs.History{{
		//	Actor:  ns.URN(InterfaceKey), // TODO??????????????????????????????????????????
		//	Key:    vcs.CreatedAction,
		//	DoneAt: time.Now(),
		//}}
		// TODO!!! add file info

		return nil
	}

	return fmt.Errorf("no data and no file path to import")
}

const onOut = "on transformTableCSV.Out(): "

func (transformOp *transformTableCSV) Out(selector *selectors.Term, params common.Map) (data structures.Pack, err error) {
	separator := params.StringDefault("separator", "")
	if separator == "" {
		return nil, fmt.Errorf(onOut + ": no 'separator' value in params")
	}

	var rowsStr []string

	if transformOp.table != nil {
		for _, row := range transformOp.table.Rows {
			// TODO!!! escape separators

			rowsStr = append(rowsStr, strings.Join(row, separator))
		}
	}

	return &structures.PackAny{
		PackDescription: transformOp.table.PackDescription,
		PackData:        structures.NewDataAny(strings.Join(rowsStr, "\n")),
	}, nil
}

const onStat = "on transformTableCSV.Stat(): "

func (transformOp *transformTableCSV) Stat(selector *selectors.Term, params common.Map) (interface{}, error) {
	return transformOp.table.Stat()
}

func (transformOp *transformTableCSV) Copy(selector *selectors.Term, params common.Map) (interface{}, error) {
	return transformOp.table, nil
}
