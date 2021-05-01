package transfer_records_types01

import (
	"fmt"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/auth"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/rbac"
	"github.com/pavlo67/common/common/starter"

	"github.com/pavlo67/data/entities/records"
)

const InterfaceKey joiner.InterfaceKey = "transfer_records_types01"

func Starter() starter.Operator {
	return &transferRecordsTypes01Starter{}
}

// ---------------------------------------------------------------------------------

var l logger.Operator
var _ starter.Operator = &transferRecordsTypes01Starter{}

type transferRecordsTypes01Starter struct {
	recordsKey   joiner.InterfaceKey
	interfaceKey joiner.InterfaceKey
}

func (tppos *transferRecordsTypes01Starter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (tppos *transferRecordsTypes01Starter) Prepare(cfg *config.Config, options common.Map) error {
	tppos.recordsKey = joiner.InterfaceKey(options.StringDefault("records_key", string(records.InterfaceKey)))
	tppos.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(InterfaceKey)))

	return nil
}

func (tppos *transferRecordsTypes01Starter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	recordsOp, _ := joinerOp.Interface(tppos.recordsKey).(records.Operator)
	if recordsOp == nil {
		return fmt.Errorf("no records.Operator with key %s", tppos.recordsKey)
	}

	transferOp, err := New(recordsOp, auth.IdentityWithRoles(rbac.RoleAdmin))
	if err != nil {
		return err
	}

	if err = joinerOp.Join(transferOp, tppos.interfaceKey); err != nil {
		return errors.CommonError(err, fmt.Sprintf("can't join *transferRecordsTypes01 as transfer.Operator with key '%s'", tppos.interfaceKey))
	}

	return nil
}
