package transform_structure_table

import (
	"fmt"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
)

const InterfaceKey joiner.InterfaceKey = "transform_structure_data_table"

func Starter() starter.Operator {
	return &transformStructureTableStarter{}
}

// ---------------------------------------------------------------------------------

var l logger.Operator
var _ starter.Operator = &transformStructureTableStarter{}

type transformStructureTableStarter struct {
	interfaceKey joiner.InterfaceKey
}

func (ttbcs *transformStructureTableStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (ttbcs *transformStructureTableStarter) Prepare(cfg *config.Config, options common.Map) error {
	ttbcs.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(InterfaceKey)))

	return nil
}

func (ttbcs *transformStructureTableStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	transformOp, err := New()
	if err != nil {
		return errors.CommonError(err, "can't init *transformStructureDataTable{} as transform.Operator")
	}

	if err = joinerOp.Join(transformOp, ttbcs.interfaceKey); err != nil {
		return errors.CommonError(err, fmt.Sprintf("can't join *transformStructureDataTable{} as transform.Operator with key '%s'", ttbcs.interfaceKey))
	}

	return nil
}
