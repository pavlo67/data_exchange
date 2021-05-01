package transfer_stdout_any

import (
	"fmt"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
)

const InterfaceKey joiner.InterfaceKey = "transfer_stdout_any"

const ModeJSON = "json"
const ModeYAML = "yaml"
const ModeTabbed = "tabbed"

func Starter() starter.Operator {
	return &transferStdoutAnyStarter{}
}

// ---------------------------------------------------------------------------------

var l logger.Operator
var _ starter.Operator = &transferStdoutAnyStarter{}

type transferStdoutAnyStarter struct {
	mode         string
	path         string
	interfaceKey joiner.InterfaceKey
}

func (tsas *transferStdoutAnyStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (tsas *transferStdoutAnyStarter) Prepare(cfg *config.Config, options common.Map) error {
	tsas.mode = options.StringDefault("mode", ModeJSON)
	tsas.path = options.StringDefault("path", "")

	tsas.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(InterfaceKey)))

	return nil
}

func (tsas *transferStdoutAnyStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	transferOp, err := New(tsas.mode, tsas.path)
	if err != nil {
		return err
	}

	if err = joinerOp.Join(transferOp, tsas.interfaceKey); err != nil {
		return errors.CommonError(err, fmt.Sprintf("can't join *transferStdoutAny{} as transfer.Operator with key '%s'", tsas.interfaceKey))
	}

	return nil
}
