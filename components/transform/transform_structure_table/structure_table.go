package transform_structure_table

import (
	"fmt"
	"strings"

	"github.com/pavlo67/common/common/errors"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data_exchange/components/transform"
)

var _ transform.Operator = &transformStructureTable{}

type transformStructureTable struct {
	data transform.Structure
}

const onNew = "on transformStructureTable.New(): "

func New() (transform.Operator, error) {
	transformOp := transformStructureTable{}
	return &transformOp, nil
}

func (transformOp *transformStructureTable) Reset() error {
	transformOp.data.Title, transformOp.data.Fields, transformOp.data.Table = "", nil, nil
	return nil
}

const onStat = "on transformStructureTable.Stat(): "

func (transformOp *transformStructureTable) Stat(params common.Map) error {
	return common.ErrNotImplemented
}

const onIn = "on transformStructureTable.In(): "

func (transformOp *transformStructureTable) In(selector *selectors.Term, data interface{}) error {
	if err := transformOp.Reset(); err != nil {
		return errors.CommonError(err, onIn)
	}

	var tablePtr *transform.Table

	switch v := data.(type) {
	case transform.Table:
		tablePtr = &v
	case *transform.Table:
		tablePtr = v
	case transform.Structure:
		tablePtr = &v.Table
		transformOp.data.Title = v.Title
	case *transform.Structure:
		if v != nil {
			tablePtr = &v.Table
			transformOp.data.Title = v.Title
		}
	}

	if tablePtr == nil || len(*tablePtr) < 1 {
		return fmt.Errorf("wrong data to import: %#v", data)
	}

	fieldIndex := make([]int, len((*tablePtr)[0]))

FIELDS:
	for i, fieldName := range (*tablePtr)[0] {
		if fieldName = strings.TrimSpace(fieldName); fieldName != "" {
			for fieldI, field := range transformOp.data.Fields {
				if field.Name == fieldName {
					fieldIndex[i] = fieldI
					continue FIELDS
				}
			}
			transformOp.data.Fields = append(transformOp.data.Fields, transform.Field{
				Name: fieldName,
			})
			fieldIndex[i] = len(transformOp.data.Fields) - 1
		} else {
			fieldIndex[i] = -1
		}
	}

	transformOp.data.Table = make(transform.Table, len(*tablePtr)-1)

	for j, row := range (*tablePtr)[1:] {
		tableRow := make([]string, len(fieldIndex))
		for i, value := range row {
			if i >= len(fieldIndex) {
				break
			}
			if fieldIndex[i] >= 0 {
				if tableRow[fieldIndex[i]] != "" {
					tableRow[fieldIndex[i]] += " " + value
				} else {
					tableRow[fieldIndex[i]] = value
				}
			}
		}

		transformOp.data.Table[j-1] = tableRow
	}

	return nil
}

func (transformOp *transformStructureTable) Out(selector *selectors.Term) (data interface{}, err error) {
	return transformOp.data, nil
}
