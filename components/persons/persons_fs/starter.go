package persons_fs

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"

	"github.com/pavlo67/data_exchange/components/persons"
)

func Starter() starter.Operator {
	return &personsFSStubStarter{}
}

const configKeyDefault = "persons_fs"

type personsFSStubStarter struct {
	interfaceKey        joiner.InterfaceKey
	interfaceCleanerKey joiner.InterfaceKey

	cfg config.Access
}

var _ starter.Operator = &personsFSStubStarter{}
var l logger.Operator

func (uks *personsFSStubStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (uks *personsFSStubStarter) Prepare(cfg *config.Config, options common.Map) error {
	configKey := options.StringDefault("config_key", configKeyDefault)

	if err := cfg.Value(configKey, &uks.cfg); err != nil {
		return err
	}

	uks.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(persons.InterfaceKey)))
	uks.interfaceCleanerKey = joiner.InterfaceKey(options.StringDefault("interface_cleaner_key", string(persons.InterfaceCleanerKey)))

	return nil
}

func (uks *personsFSStubStarter) Run(joinerOp joiner.Operator) error {

	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	personsOp, personsCleanerOp, err := New(uks.cfg)
	if err != nil {
		return errors.Wrapf(err, "can't personsFSStub.New()")
	}

	if err = joinerOp.Join(personsOp, uks.interfaceKey); err != nil {
		return errors.Wrap(err, "can't join *personsFSStub{} as persons.Operator interface")
	}

	if err = joinerOp.Join(personsCleanerOp, uks.interfaceCleanerKey); err != nil {
		return errors.Wrap(err, "can't join *personsFSStub{} as db.Cleaner interface")
	}

	return nil
}
