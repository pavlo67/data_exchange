package exchange_0_1_files

import (
	"fmt"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
)

func Starter() starter.Operator {
	return &exchange01FilesStarter{}
}

var l logger.Operator
var _ starter.Operator = &exchange01FilesStarter{}

type exchange01FilesStarter struct {
	path string

	interfaceKey joiner.InterfaceKey
	cleanerKey   joiner.InterfaceKey
}

func (e01fs *exchange01FilesStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (e01fs *exchange01FilesStarter) Prepare(cfg *config.Config, options common.Map) error {
	e01fs.path = options.StringDefault("path", "")

	e01fs.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(InterfaceKey)))
	e01fs.cleanerKey = joiner.InterfaceKey(options.StringDefault("cleaner_key", string(InterfaceKeyCleaner)))

	return nil
}

func (e01fs *exchange01FilesStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	exchange01FilesOp, exchange01FilesCleanerOp, err := New(e01fs.path)
	if err != nil {
		return errors.Wrap(err, "can't init *exchange01Files{} as exchange.Operator")
	}

	if err = joinerOp.Join(exchange01FilesOp, e01fs.interfaceKey); err != nil {
		return errors.Wrapf(err, "can't join *exchange01Files{} as exchange.Operator with key '%s'", e01fs.interfaceKey)
	}

	if err = joinerOp.Join(exchange01FilesCleanerOp, e01fs.cleanerKey); err != nil {
		return errors.Wrapf(err, "can't join *exchange01Files{} as crud.Cleaner with key '%s'", e01fs.cleanerKey)
	}

	return nil
}
