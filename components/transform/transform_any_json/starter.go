package transform_any_json

import (
	"fmt"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
)

const InterfaceKey joiner.InterfaceKey = "transform_any_json"

func Starter() starter.Operator {
	return &transformAnyJSONStarter{}
}

// ---------------------------------------------------------------------------------

var l logger.Operator
var _ starter.Operator = &transformAnyJSONStarter{}

type transformAnyJSONStarter struct {
	path           string
	exemplar       interface{}
	prefix, indent string
	interfaceKey   joiner.InterfaceKey
}

func (tajs *transformAnyJSONStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (tajs *transformAnyJSONStarter) Prepare(cfg *config.Config, options common.Map) error {
	tajs.path = options.StringDefault("path", "")
	tajs.exemplar = options["exemplar"]
	tajs.prefix = options.StringDefault("prefix", "")
	tajs.indent = options.StringDefault("indent", "")

	tajs.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(InterfaceKey)))

	return nil
}

func (tajs *transformAnyJSONStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	transformOp, err := New(tajs.path, tajs.exemplar, tajs.prefix, tajs.indent)
	if err != nil {
		return errors.CommonError(err, "can't init *transformAnyJSON{} as transform.Operator")
	}

	if err = joinerOp.Join(transformOp, tajs.interfaceKey); err != nil {
		return errors.CommonError(err, fmt.Sprintf("can't join *transformAnyJSON{} as transform.Operator with key '%s'", tajs.interfaceKey))
	}

	return nil
}
