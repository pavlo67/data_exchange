package transfer_table_csv

import (
	"fmt"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
)

const InterfaceKey joiner.InterfaceKey = "transfer_table_csv"

func Starter() starter.Operator {
	return &transferTableCSVStarter{}
}

// ---------------------------------------------------------------------------------

var l logger.Operator
var _ starter.Operator = &transferTableCSVStarter{}

type transferTableCSVStarter struct {
	interfaceKey joiner.InterfaceKey
}

func (ttcs *transferTableCSVStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (ttcs *transferTableCSVStarter) Prepare(cfg *config.Config, options common.Map) error {
	ttcs.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(InterfaceKey)))

	return nil
}

func (ttcs *transferTableCSVStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	transferOp, err := New()
	if err != nil {
		return err
	}

	if err = joinerOp.Join(transferOp, ttcs.interfaceKey); err != nil {
		return errors.CommonError(err, fmt.Sprintf("can't join *transferTableCSV{} as transfer.Operator with key '%s'", ttcs.interfaceKey))
	}

	return nil
}
