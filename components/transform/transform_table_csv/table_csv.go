package transform_table_csv

import (
	"fmt"
	"strings"
	"time"

	"github.com/pavlo67/data_exchange/components/ns"
	"github.com/pavlo67/data_exchange/components/vcs"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"
	"github.com/pavlo67/data_exchange/components/extraction"

	"github.com/pavlo67/data_exchange/components/transform"
)

var _ transform.Operator = &transformTableCSV{}

type transformTableCSV struct {
	// TODO!!! use .Structure

	title   string
	table   transform.Table
	history vcs.History

	path      string
	separator string
}

const onNew = "on transformTableCSV.New(): "

func New(path, separator string) (transform.Operator, error) {

	if separator == "" {
		return nil, fmt.Errorf("no 'separator' value")
	}
	transformOp := transformTableCSV{
		path:      strings.TrimSpace(path),
		separator: separator,
	}
	return &transformOp, nil
}

func (transformOp *transformTableCSV) Reset() error {
	transformOp.title, transformOp.table, transformOp.history = "", nil, nil
	return nil
}

const onStat = "on transformTableCSV.Stat(): "

func (transformOp *transformTableCSV) Stat(params common.Map) error {
	return common.ErrNotImplemented
}

const onIn = "on transformTableCSV.In(): "

func (transformOp *transformTableCSV) In(selector *selectors.Term, data interface{}) error {
	if err := transformOp.Reset(); err != nil {
		return errors.CommonError(err, onIn)
	}

	var err error

	if data != nil {
		if bytes, _ := data.([]byte); bytes != nil {
			transformOp.table, err = extraction.TableBytes(bytes, transformOp.separator)
			if err != nil {
				return errors.CommonError(err, onIn)
			}
			transformOp.history = vcs.History{{
				Actor:  ns.ID(InterfaceKey),
				Key:    vcs.CreatedAction,
				DoneAt: time.Now(),
			}}
			return nil
		}

		return fmt.Errorf("wrong data to import: %#v", data)
	}

	if transformOp.path != "" {
		_, transformOp.table, err = extraction.TableFile(transformOp.path, transformOp.separator)
		if err != nil {
			return errors.CommonError(err, onIn)
		}
		transformOp.title = transformOp.path
		transformOp.history = vcs.History{{
			Actor:  ns.ID(InterfaceKey),
			Key:    vcs.CreatedAction,
			DoneAt: time.Now(),
		}}
		// TODO!!! add file info

		return nil
	}

	return fmt.Errorf("no data and no file path to import")
}

func (transformOp *transformTableCSV) Out(selector *selectors.Term) (data interface{}, err error) {
	return transform.Structure{
		Title:   transformOp.title,
		Fields:  nil,
		Table:   transformOp.table,
		History: transformOp.history,
	}, nil
}
