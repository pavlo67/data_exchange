package transformer_stdout_any

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/pavlo67/data_exchange/components/structures"
	"github.com/pavlo67/data_exchange/components/transformer"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"
)

var _ transformer.Operator = &transformStdoutAny{}

type transformStdoutAny struct {
	any  interface{}
	mode string
	path string
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

func (transformOp *transformStdoutAny) Name() string {
	return string(InterfaceKey)
}

func (transformOp *transformStdoutAny) Reset() error {
	transformOp.any = nil
	return nil
}

const onStat = "on transformStdoutAny.Stat(): "

func (transformOp *transformStdoutAny) Stat(selector *selectors.Term, params common.Map) (interface{}, error) {
	return nil, common.ErrNotImplemented
}

const onIn = "on transformStdoutAny.In(): "

// DEPRECATED: use common.ErrNotSupported
var ErrNotSupported = errors.New("not_supported")

func (transformOp *transformStdoutAny) In(selector *selectors.Term, params common.Map, data interface{}) error {
	transformOp.any = data
	return nil
}

func (transformOp *transformStdoutAny) Out(selector *selectors.Term, params common.Map) (data interface{}, err error) {
	return transformOp.any, nil
}

const onCopy = "on transformStdoutAny.Copy(): "

func (transformOp *transformStdoutAny) Copy(selector *selectors.Term, params common.Map) (data interface{}, err error) {

	startLines := int(params.Int64Default("start_lines", 25))
	endLines := int(params.Int64Default("end_lines", 25))

	var items []interface{}

	switch v := transformOp.any.(type) {
	case structures.Table:
		for _, line := range v.Data {
			items = append(items, line)
		}
	case *structures.Table:
		if v == nil {
			return nil, fmt.Errorf(onCopy+": nil data (%T)", transformOp.any)
		}
		for _, line := range v.Data {
			items = append(items, line)
		}
	case structures.Pack:
		if v.Data == nil {
			return nil, fmt.Errorf(onCopy+": nil data (%T)", transformOp.any)
		}
		data := reflect.ValueOf(v.Data)

		for i := 0; i < data.Len(); i++ {
			items = append(items, data.Index(i).Interface())
		}
	case *structures.Pack:
		if v == nil || v.Data == nil {
			return nil, fmt.Errorf(onCopy+": nil data (%T)", transformOp.any)
		}
		data := reflect.ValueOf(v.Data)
		for i := 0; i < data.Len(); i++ {
			items = append(items, data.Index(i).Interface())
		}
	case []interface{}:
		items = v
	case *[]interface{}:
		if v == nil {
			return nil, fmt.Errorf(onCopy+": nil data (%T)", transformOp.any)
		}
		items = *v
	default:
		return nil, fmt.Errorf(onCopy+": wrong data type (%T)", transformOp.any)
	}

	var dataToOut []interface{}
	if startLines < 0 || startLines >= len(items) || startLines+endLines >= len(items) {
		dataToOut = items
	} else {
		dataToOut = append(items[:startLines], items[len(items)-endLines:]...)
	}

	mode := params.StringDefault("mode", transformOp.mode)

	var bytes []byte

	if transformOp.any != nil {
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
