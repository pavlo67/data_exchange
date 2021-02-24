package exchange_file_tabbed

import (
	"fmt"

	"github.com/pavlo67/data_exchange/components/exchange"

	"github.com/pkg/errors"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
)

func Starter() starter.Operator {
	return &exchangeFileTabbedStarter{}
}

var l logger.Operator
var _ starter.Operator = &exchangeFileTabbedStarter{}

type exchangeFileTabbedStarter struct {
	path string

	interfaceKey joiner.InterfaceKey
}

func (efts *exchangeFileTabbedStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (efts *exchangeFileTabbedStarter) Prepare(cfg *config.Config, options common.Map) error {
	efts.path = options.StringDefault("path", string(exchange.InterfaceKey))
	efts.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(exchange.InterfaceKey)))

	return nil
}

func (efts *exchangeFileTabbedStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	exytractorOp, err := New(efts.path)
	if err != nil {
		return errors.Wrap(err, "can't init *exchangeFileTabbed{} as exchange.Operator")
	}

	if err = joinerOp.Join(exytractorOp, efts.interfaceKey); err != nil {
		return errors.Wrapf(err, "can't join *exchangeFileTabbed{} as exchange.Operator with key '%s'", efts.interfaceKey)
	}

	return nil
}
