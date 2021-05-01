package transfer_persons_types01

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

	"github.com/pavlo67/data/entities/persons"
)

const InterfaceKey joiner.InterfaceKey = "transfer_operator_persons_pack"

func Starter() starter.Operator {
	return &transferPersonsOperatorPackTypes01Starter{}
}

// ---------------------------------------------------------------------------------

var l logger.Operator
var _ starter.Operator = &transferPersonsOperatorPackTypes01Starter{}

type transferPersonsOperatorPackTypes01Starter struct {
	personsKey   joiner.InterfaceKey
	interfaceKey joiner.InterfaceKey
}

func (tppos *transferPersonsOperatorPackTypes01Starter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (tppos *transferPersonsOperatorPackTypes01Starter) Prepare(cfg *config.Config, options common.Map) error {
	tppos.personsKey = joiner.InterfaceKey(options.StringDefault("persons_key", string(persons.InterfaceKey)))
	tppos.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(InterfaceKey)))

	return nil
}

func (tppos *transferPersonsOperatorPackTypes01Starter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	personsOp, _ := joinerOp.Interface(tppos.personsKey).(persons.Operator)
	if personsOp == nil {
		return fmt.Errorf("no persons.Operator with key %s", tppos.personsKey)
	}

	transferOp, err := New(personsOp, auth.IdentityWithRoles(rbac.RoleAdmin))
	if err != nil {
		return err
	}

	if err = joinerOp.Join(transferOp, tppos.interfaceKey); err != nil {
		return errors.CommonError(err, fmt.Sprintf("can't join *transferPersonsTypes01 as transfer.Operator with key '%s'", tppos.interfaceKey))
	}

	return nil
}
