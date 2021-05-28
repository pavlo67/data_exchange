package transfer_json_any

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pavlo67/data/components/structures"

	"github.com/pavlo67/common/common/errors"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/selectors"
	"github.com/pavlo67/data/components/transfer"
)

var _ transfer.Operator = &transferJSONAny{}

type transferJSONAny struct {
	packAny *structures.PackAny
}

const onNew = "on transferJSONAny.New(): "

func New(path string, exemplar interface{}, prefix, indent string) (transfer.Operator, error) {
	transferOp := transferJSONAny{}
	return &transferOp, nil
}

func (transferOp *transferJSONAny) reset() error {
	transferOp.packAny = nil
	return nil
}

func (transferOp *transferJSONAny) Name() string {
	return string(InterfaceKey)
}

const onIn = "on transferJSONAny.In(): "

func (transferOp *transferJSONAny) In(pack structures.Pack, params common.Map) error {
	if pack == nil {
		return errors.New(onIn + "nil pack to import")
	}

	transferOp.packAny = &structures.PackAny{
		ItemDescription: pack.Description(),
		PackData:        structures.NewDataAny(pack.Data().Value()),
	}
	return nil
}

func (transferOp *transferJSONAny) Out(selector *selectors.Term, params common.Map) (pack structures.Pack, err error) {
	return transferOp.packAny, nil
}

const onStat = "on transferJSONAny.Stat(): "

func (transferOp *transferJSONAny) Stat(selector *selectors.Term, params common.Map) (interface{}, error) {
	// TODO!!! type switch or get Stat interface

	return nil, common.ErrNotImplemented
}

const onCopy = "on transferJSONAny.Copy(): "

func (transferOp *transferJSONAny) Copy(selector *selectors.Term, params common.Map) (interface{}, error) {
	prefix := params.StringDefault("prefix", "")
	indent := params.StringDefault("indent", "")
	path := strings.TrimSpace(params.StringDefault("path", ""))

	var bytes []byte
	var err error

	if transferOp.packAny != nil {
		if prefix+indent != "" {
			bytes, err = json.MarshalIndent(transferOp.packAny, prefix, indent)
		} else {
			bytes, err = json.Marshal(transferOp.packAny)
		}
		if err != nil {
			return nil, fmt.Errorf(onCopy+": marshalling %#v got %s", transferOp.packAny, err)
		}
	}

	if path != "" {
		if err = ioutil.WriteFile(path, bytes, 644); err != nil {
			return nil, fmt.Errorf(onCopy+": writing %s got %s", path, err)
		}
	}

	return bytes, nil
}
