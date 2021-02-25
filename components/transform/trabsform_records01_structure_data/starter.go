package transform_records_01

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
	return &transform01FilesStarter{}
}

var l logger.Operator
var _ starter.Operator = &transform01FilesStarter{}

type transform01FilesStarter struct {
	path string

	interfaceKey joiner.InterfaceKey
	cleanerKey   joiner.InterfaceKey
}

func (e01fs *transform01FilesStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (e01fs *transform01FilesStarter) Prepare(cfg *config.Config, options common.Map) error {
	e01fs.path = options.StringDefault("path", "")

	e01fs.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(InterfaceKey)))
	e01fs.cleanerKey = joiner.InterfaceKey(options.StringDefault("cleaner_key", string(InterfaceKeyCleaner)))

	return nil
}

func (e01fs *transform01FilesStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	transform01FilesOp, transform01FilesCleanerOp, err := New(e01fs.path)
	if err != nil {
		return errors.Wrap(err, "can't init *transform01Files{} as transform.Operator")
	}

	if err = joinerOp.Join(transform01FilesOp, e01fs.interfaceKey); err != nil {
		return errors.Wrapf(err, "can't join *transform01Files{} as transform.Operator with key '%s'", e01fs.interfaceKey)
	}

	if err = joinerOp.Join(transform01FilesCleanerOp, e01fs.cleanerKey); err != nil {
		return errors.Wrapf(err, "can't join *transform01Files{} as crud.Cleaner with key '%s'", e01fs.cleanerKey)
	}

	return nil
}
