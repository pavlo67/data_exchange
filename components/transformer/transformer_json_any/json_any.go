package transformer_json_any

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"
	"github.com/pavlo67/data_exchange/components/transformer"
)

var _ transformer.Operator = &transformerJSONAny{}

type transformerJSONAny struct {
	any interface{}
}

const onNew = "on transformerJSONAny.New(): "

func New(path string, exemplar interface{}, prefix, indent string) (transformer.Operator, error) {
	transformOp := transformerJSONAny{}
	return &transformOp, nil
}

func (transformOp *transformerJSONAny) Name() string {
	return string(InterfaceKey)
}

func (transformOp *transformerJSONAny) Reset() error {
	transformOp.any = nil
	return nil
}

const onStat = "on transformerJSONAny.Stat(): "

func (transformOp *transformerJSONAny) Stat(selector *selectors.Term, params common.Map) (interface{}, error) {
	// TODO!!! type switch or get Stat interface

	return nil, common.ErrNotImplemented
}

func (transformOp *transformerJSONAny) In(selector *selectors.Term, params common.Map, data interface{}) error {
	transformOp.any = data
	return nil
}

func (transformOp *transformerJSONAny) Out(selector *selectors.Term, params common.Map) (data interface{}, err error) {
	return transformOp.any, nil
}

const onCopy = "on transformerJSONAny.Copy(): "

func (transformOp *transformerJSONAny) Copy(selector *selectors.Term, params common.Map) (data interface{}, err error) {
	prefix := params.StringDefault("prefix", "")
	indent := params.StringDefault("indent", "")
	path := strings.TrimSpace(params.StringDefault("path", ""))

	var bytes []byte

	if transformOp.any != nil {
		if prefix+indent != "" {
			bytes, err = json.MarshalIndent(transformOp.any, prefix, indent)
		} else {
			bytes, err = json.Marshal(transformOp.any)
		}
		if err != nil {
			return nil, fmt.Errorf(onCopy+": marshalling %#v got %s", transformOp.any, err)
		}
	}

	if path != "" {
		if err = ioutil.WriteFile(path, bytes, 644); err != nil {
			return nil, fmt.Errorf(onCopy+": writing %s got %s", path, err)
		}
	}

	return bytes, nil
}
