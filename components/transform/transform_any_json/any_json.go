package transform_any_json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data_exchange/components/transform"
)

var _ transform.Operator = &transformAnyJSON{}

type transformAnyJSON struct {
	any            interface{}
	path           string
	exemplar       interface{}
	prefix, indent string
}

const onNew = "on transformAnyJSON.New(): "

func New(path string, exemplar interface{}, prefix, indent string) (transform.Operator, error) {
	transformOp := transformAnyJSON{
		path:     strings.TrimSpace(path),
		exemplar: exemplar,
		prefix:   prefix,
		indent:   indent,
	}
	return &transformOp, nil
}

func (transformOp *transformAnyJSON) Reset() error {
	transformOp.any = nil
	return nil
}

const onStat = "on transformAnyJSON.Stat(): "

func (transformOp *transformAnyJSON) Stat(params common.Map) error {
	return common.ErrNotImplemented
}

const onIn = "on transformAnyJSON.In(): "

func (transformOp *transformAnyJSON) In(selector *selectors.Term, data interface{}) error {
	if err := transformOp.Reset(); err != nil {
		return errors.CommonError(err, onIn)
	}

	var err error
	var bytes []byte

	if data != nil && transformOp.exemplar != nil {
		if bytes, _ = data.([]byte); bytes == nil {
			return fmt.Errorf(onIn+": wrong data to import: %#v", data)
		}
	}

	if transformOp.path != "" {
		if bytes, err = ioutil.ReadFile(transformOp.path); err != nil {
			return fmt.Errorf(onIn+": reading %s got %s", transformOp.path, err)
		}
	}

	if bytes == nil {
		return errors.New(onIn + ": no bytes data and no file path to import")
	}

	if len(bytes) < 1 {
		transformOp.any = nil
		return nil
	}

	if err = json.Unmarshal(bytes, &transformOp.exemplar); err != nil {
		return errors.CommonError(err, fmt.Sprintf(onIn+": unmarshalling %s into %T", bytes, transformOp.exemplar))
	}

	transformOp.any = transformOp.exemplar
	return nil
}

const onOut = "on transformAnyJSON.Out(): "

func (transformOp *transformAnyJSON) Out(selector *selectors.Term) (data interface{}, err error) {

	var bytes []byte

	if transformOp.any != nil {
		if transformOp.prefix+transformOp.indent != "" {
			bytes, err = json.MarshalIndent(transformOp.any, transformOp.prefix, transformOp.indent)
		} else {
			bytes, err = json.Marshal(transformOp.any)
		}
		if err != nil {
			return nil, fmt.Errorf(onOut+": marshalling %#v got %s", transformOp.any, err)
		}
	}

	if transformOp.path != "" {
		if err = ioutil.WriteFile(transformOp.path, bytes, 644); err != nil {
			return nil, fmt.Errorf(onOut+": writing %s got %s", transformOp.path, err)
		}
	}

	return bytes, nil
}
