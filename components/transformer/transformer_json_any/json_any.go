package transformer_json_any

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pavlo67/data_exchange/components/structures"

	"github.com/pavlo67/common/common/errors"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"
	"github.com/pavlo67/data_exchange/components/transformer"
)

var _ transformer.Operator = &transformerJSONAny{}

type transformerJSONAny struct {
	packAny *structures.PackAny
}

const onNew = "on transformerJSONAny.New(): "

func New(path string, exemplar interface{}, prefix, indent string) (transformer.Operator, error) {
	transformOp := transformerJSONAny{}
	return &transformOp, nil
}

func (transformOp *transformerJSONAny) reset() error {
	transformOp.packAny = nil
	return nil
}

func (transformOp *transformerJSONAny) Name() string {
	return string(InterfaceKey)
}

const onIn = "on transformerJSONAny.In(): "

func (transformOp *transformerJSONAny) In(pack structures.Pack, params common.Map) error {
	if pack == nil {
		return errors.New(onIn + "nil pack to import")
	}

	transformOp.packAny = &structures.PackAny{
		PackDescription: pack.Description(),
		PackData:        pack.Data(),
	}
	return nil
}

func (transformOp *transformerJSONAny) Out(selector *selectors.Term, params common.Map) (pack structures.Pack, err error) {
	return transformOp.packAny, nil
}

const onStat = "on transformerJSONAny.Stat(): "

func (transformOp *transformerJSONAny) Stat(selector *selectors.Term, params common.Map) (interface{}, error) {
	// TODO!!! type switch or get Stat interface

	return nil, common.ErrNotImplemented
}

const onCopy = "on transformerJSONAny.Copy(): "

func (transformOp *transformerJSONAny) Copy(selector *selectors.Term, params common.Map) (interface{}, error) {
	prefix := params.StringDefault("prefix", "")
	indent := params.StringDefault("indent", "")
	path := strings.TrimSpace(params.StringDefault("path", ""))

	var bytes []byte
	var err error

	if transformOp.packAny != nil {
		if prefix+indent != "" {
			bytes, err = json.MarshalIndent(transformOp.packAny, prefix, indent)
		} else {
			bytes, err = json.Marshal(transformOp.packAny)
		}
		if err != nil {
			return nil, fmt.Errorf(onCopy+": marshalling %#v got %s", transformOp.packAny, err)
		}
	}

	if path != "" {
		if err = ioutil.WriteFile(path, bytes, 644); err != nil {
			return nil, fmt.Errorf(onCopy+": writing %s got %s", path, err)
		}
	}

	return bytes, nil
}
