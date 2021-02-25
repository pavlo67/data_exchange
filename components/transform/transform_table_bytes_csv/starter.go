package transform_table_bytes_csv

import (
	"fmt"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
)

const InterfaceKey joiner.InterfaceKey = "transform_table_bytes_csv"

func Starter() starter.Operator {
	return &transformTableBytesCSVStarter{}
}

// ---------------------------------------------------------------------------------

var l logger.Operator
var _ starter.Operator = &transformTableBytesCSVStarter{}

type transformTableBytesCSVStarter struct {
	path         string
	separator    string
	interfaceKey joiner.InterfaceKey
}

func (ttbcs *transformTableBytesCSVStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (ttbcs *transformTableBytesCSVStarter) Prepare(cfg *config.Config, options common.Map) error {
	ttbcs.path = options.StringDefault("path", "")
	if ttbcs.separator = options.StringDefault("separator", ""); ttbcs.separator == "" {
		return fmt.Errorf("no 'separator' value in options: %#v", options)
	}

	ttbcs.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(InterfaceKey)))

	return nil
}

func (ttbcs *transformTableBytesCSVStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	transformOp, err := New(ttbcs.path, ttbcs.separator)
	if err != nil {
		return errors.CommonError(err, "can't init *transformStructureDataTable{} as transform.Operator")
	}

	if err = joinerOp.Join(transformOp, ttbcs.interfaceKey); err != nil {
		return errors.CommonError(err, fmt.Sprintf("can't join *transformStructureDataTable{} as transform.Operator with key '%s'", ttbcs.interfaceKey))
	}

	return nil
}
