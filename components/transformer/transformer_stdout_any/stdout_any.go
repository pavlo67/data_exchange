package transformer_stdout_any

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"

	"github.com/pavlo67/data/components/structures"
	"github.com/pavlo67/data/components/transformer"
)

var _ transformer.Operator = &transformStdoutAny{}

type transformStdoutAny struct {
	packAny *structures.PackAny
	mode    string
	path    string
}

const onNew = "on transformStdoutAny.New(): "

func New(mode, path string) (transformer.Operator, error) {
	if mode = strings.TrimSpace(mode); mode == "" {
		mode = ModeJSON
	}

	transformOp := transformStdoutAny{
		mode: mode,
		path: strings.TrimSpace(path),
	}
	return &transformOp, nil
}

func (transformOp *transformStdoutAny) reset() error {
	transformOp.packAny = nil
	return nil
}

func (transformOp *transformStdoutAny) Name() string {
	return string(InterfaceKey)
}

//In(pack Pack, params common.Map) error                                   // import from external source
//Out(selector *selectors.Term, params common.Map) (pack Pack, err error)  // export to external source
//Stat(selector *selectors.Term, params common.Map) (pack Pack, err error) // internal storage statistics
//Copy(selector *selectors.Term, params common.Map) (interface{}, error)   // internal storage snapshot

const onIn = "on transformStdoutAny.In(): "

// DEPRECATED: use common.ErrNotSupported
var ErrNotSupported = errors.New("not_supported")

func (transformOp *transformStdoutAny) In(pack structures.Pack, params common.Map) error {
	if pack == nil {
		return errors.New(onIn + "nil pack to import")
	}

	transformOp.packAny = &structures.PackAny{
		PackDescription: pack.Description(),
		PackData:        structures.NewDataAny(pack.Data().Value()),
	}
	return nil
}

func (transformOp *transformStdoutAny) Out(selector *selectors.Term, params common.Map) (structures.Pack, error) {
	return transformOp.packAny, nil
}

const onStat = "on transformStdoutAny.Stat(): "

func (transformOp *transformStdoutAny) Stat(selector *selectors.Term, params common.Map) (interface{}, error) {
	return nil, common.ErrNotImplemented
}

const onCopy = "on transformStdoutAny.Copy(): "

func (transformOp *transformStdoutAny) Copy(selector *selectors.Term, params common.Map) (interface{}, error) {

	startLines := int(params.Int64Default("start_lines", 25))
	endLines := int(params.Int64Default("end_lines", 25))

	var items []interface{}

	switch v := transformOp.packAny.PackData.Value().(type) {
	case structures.Rows:
		for _, line := range v {
			items = append(items, line)
		}
	case *structures.Rows:
		if v == nil {
			return nil, fmt.Errorf(onCopy+": nil data (%T)", transformOp.packAny)
		}
		for _, line := range *v {
			items = append(items, line)
		}
	case []interface{}:
		items = v
	case *[]interface{}:
		if v == nil {
			return nil, fmt.Errorf(onCopy+": nil data (%T)", transformOp.packAny)
		}
		items = *v
	default:
		return nil, fmt.Errorf(onCopy+": wrong data type (%T)", transformOp.packAny)
	}

	var dataToOut []interface{}
	if startLines < 0 || startLines >= len(items) || startLines+endLines >= len(items) {
		dataToOut = items
	} else {
		dataToOut = append(items[:startLines], items[len(items)-endLines:]...)
	}

	mode := params.StringDefault("mode", transformOp.mode)

	var bytes []byte
	var err error

	if transformOp.packAny != nil {
		for _, line := range dataToOut {

			var lineBytes []byte
			switch mode {
			case ModeTabbed:
				if v, _ := line.([]string); v != nil {
					lineBytes = []byte(strings.Join(v, "\t") + "\n")
				} else {
					return nil, fmt.Errorf(onCopy+": on strings.Join(%#v): wrong line type", line)
				}
			case ModeYAML:
				lineBytes, err = yaml.Marshal(line)
				if err != nil {
					return nil, fmt.Errorf(onCopy+": on yaml.Marshal(%#v) got %s", line, err)
				}
				lineBytes = append(lineBytes, '\n', '\n')
			// case ModeJSON:
			default:
				lineBytes, err = json.MarshalIndent(line, "", "   ")
				if err != nil {
					return nil, fmt.Errorf(onCopy+": on json.Marshal(%#v) got %s", line, err)
				}
				lineBytes = append(lineBytes, '\n', '\n')
			}

			bytes = append(bytes, lineBytes...)
		}
	}

	if transformOp.path != "" {
		if err = ioutil.WriteFile(transformOp.path, bytes, 644); err != nil {
			return nil, fmt.Errorf(onCopy+": writing %s got %s", transformOp.path, err)
		}
	}

	return bytes, nil
}
