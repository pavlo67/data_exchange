package transformer_stdout_any

import (
	"fmt"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
)

const InterfaceKey joiner.InterfaceKey = "transform_stdout_any"

const ModeJSON = "json"
const ModeYAML = "yaml"
const ModeTabbed = "tabbed"

func Starter() starter.Operator {
	return &transformStdoutAnyStarter{}
}

// ---------------------------------------------------------------------------------

var l logger.Operator
var _ starter.Operator = &transformStdoutAnyStarter{}

type transformStdoutAnyStarter struct {
	mode         string
	path         string
	interfaceKey joiner.InterfaceKey
}

func (tsas *transformStdoutAnyStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (tsas *transformStdoutAnyStarter) Prepare(cfg *config.Config, options common.Map) error {
	tsas.mode = options.StringDefault("mode", ModeJSON)
	tsas.path = options.StringDefault("path", "")

	tsas.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(InterfaceKey)))

	return nil
}

func (tsas *transformStdoutAnyStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	transformOp, err := New(tsas.mode, tsas.path)
	if err != nil {
		return err
	}

	if err = joinerOp.Join(transformOp, tsas.interfaceKey); err != nil {
		return errors.CommonError(err, fmt.Sprintf("can't join *transformStdoutAny{} as transform.Operator with key '%s'", tsas.interfaceKey))
	}

	return nil
}
