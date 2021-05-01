package transfer_json_any

import (
	"fmt"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
)

const InterfaceKey joiner.InterfaceKey = "transfer_json_any"

func Starter() starter.Operator {
	return &transferJSONAnyStarter{}
}

// ---------------------------------------------------------------------------------

var l logger.Operator
var _ starter.Operator = &transferJSONAnyStarter{}

type transferJSONAnyStarter struct {
	path           string
	exemplar       interface{}
	prefix, indent string
	interfaceKey   joiner.InterfaceKey
}

func (tjas *transferJSONAnyStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (tjas *transferJSONAnyStarter) Prepare(cfg *config.Config, options common.Map) error {
	tjas.path = options.StringDefault("path", "")
	tjas.exemplar = options["exemplar"]
	tjas.prefix = options.StringDefault("prefix", "")
	tjas.indent = options.StringDefault("indent", "")

	tjas.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(InterfaceKey)))

	return nil
}

func (tjas *transferJSONAnyStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	transferOp, err := New(tjas.path, tjas.exemplar, tjas.prefix, tjas.indent)
	if err != nil {
		return err
	}

	if err = joinerOp.Join(transferOp, tjas.interfaceKey); err != nil {
		return errors.CommonError(err, fmt.Sprintf("can't join *transferJSONAny{} as transfer.Operator with key '%s'", tjas.interfaceKey))
	}

	return nil
}
