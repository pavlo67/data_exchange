package transfer_stdout_any

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
	"github.com/pavlo67/data/components/transfer"
)

var _ transfer.Operator = &transferStdoutAny{}

type transferStdoutAny struct {
	packAny *structures.PackAny
	mode    string
	path    string
}

const onNew = "on transferStdoutAny.New(): "

func New(mode, path string) (transfer.Operator, error) {
	if mode = strings.TrimSpace(mode); mode == "" {
		mode = ModeJSON
	}

	transferOp := transferStdoutAny{
		mode: mode,
		path: strings.TrimSpace(path),
	}
	return &transferOp, nil
}

func (transferOp *transferStdoutAny) reset() error {
	transferOp.packAny = nil
	return nil
}

func (transferOp *transferStdoutAny) Name() string {
	return string(InterfaceKey)
}

//In(pack Pack, params common.Map) error                                   // import from external source
//Out(selector *selectors.Term, params common.Map) (pack Pack, err error)  // export to external source
//Stat(selector *selectors.Term, params common.Map) (pack Pack, err error) // internal storage statistics
//Copy(selector *selectors.Term, params common.Map) (interface{}, error)   // internal storage snapshot

const onIn = "on transferStdoutAny.In(): "

// DEPRECATED: use common.ErrNotSupported
var ErrNotSupported = errors.New("not_supported")

func (transferOp *transferStdoutAny) In(pack structures.Pack, params common.Map) error {
	if pack == nil {
		return errors.New(onIn + "nil pack to import")
	}

	transferOp.packAny = &structures.PackAny{
		ItemDescription: pack.Description(),
		PackData:        structures.NewDataAny(pack.Data().Value()),
	}
	return nil
}

func (transferOp *transferStdoutAny) Out(selector *selectors.Term, params common.Map) (structures.Pack, error) {
	return transferOp.packAny, nil
}

const onStat = "on transferStdoutAny.Stat(): "

func (transferOp *transferStdoutAny) Stat(selector *selectors.Term, params common.Map) (interface{}, error) {
	return nil, common.ErrNotImplemented
}

const onCopy = "on transferStdoutAny.Copy(): "

func (transferOp *transferStdoutAny) Copy(selector *selectors.Term, params common.Map) (interface{}, error) {

	startLines := int(params.Int64Default("start_lines", 25))
	endLines := int(params.Int64Default("end_lines", 25))

	var items []interface{}

	switch v := transferOp.packAny.PackData.Value().(type) {
	//case structures.Rows:
	//	for _, line := range v {
	//		items = append(items, line)
	//	}
	//case *structures.Rows:
	//	if v == nil {
	//		return nil, fmt.Errorf(onCopy+": nil data (%T)", transferOp.packAny)
	//	}
	//	for _, line := range *v {
	//		items = append(items, line)
	//	}
	case []interface{}:
		items = v
	case *[]interface{}:
		if v == nil {
			return nil, fmt.Errorf(onCopy+": nil data (%T)", transferOp.packAny)
		}
		items = *v
	default:
		return nil, fmt.Errorf(onCopy+": wrong data type (%T)", transferOp.packAny)
	}

	var dataToOut []interface{}
	if startLines < 0 || startLines >= len(items) || startLines+endLines >= len(items) {
		dataToOut = items
	} else {
		dataToOut = append(items[:startLines], items[len(items)-endLines:]...)
	}

	mode := params.StringDefault("mode", transferOp.mode)

	var bytes []byte
	var err error

	if transferOp.packAny != nil {
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

	if transferOp.path != "" {
		if err = ioutil.WriteFile(transferOp.path, bytes, 644); err != nil {
			return nil, fmt.Errorf(onCopy+": writing %s got %s", transferOp.path, err)
		}
	}

	return bytes, nil
}
