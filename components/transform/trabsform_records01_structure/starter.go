package transform_records01_structure

import (
	"fmt"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
)

const InterfaceKey = joiner.InterfaceKey("transform_records01_structure")

func Starter() starter.Operator {
	return &transformRecords01StructureStarter{}
}

var l logger.Operator
var _ starter.Operator = &transformRecords01StructureStarter{}

type transformRecords01StructureStarter struct {
	interfaceKey joiner.InterfaceKey
}

func (tr01ss *transformRecords01StructureStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (tr01ss *transformRecords01StructureStarter) Prepare(cfg *config.Config, options common.Map) error {
	tr01ss.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(InterfaceKey)))

	return nil
}

func (tr01ss *transformRecords01StructureStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	transformOp, err := New()
	if err != nil {
		return errors.CommonError(err, "can't init *transformRecords01Structure{} as transform.Operator")
	}

	if err = joinerOp.Join(transformOp, tr01ss.interfaceKey); err != nil {
		return errors.CommonError(err, fmt.Sprintf("can't join *transformRecords01Structure{} as transform.Operator with key '%s'", tr01ss.interfaceKey))
	}

	return nil
}
