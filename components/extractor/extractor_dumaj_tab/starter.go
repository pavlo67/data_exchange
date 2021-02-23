package extractor_dumaj_tab

import (
	"fmt"

	"github.com/pavlo67/data_exchange/components/extractor"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
	"github.com/pkg/errors"
)

func Starter() starter.Operator {
	return &extractorDumajTabStarter{}
}

var l logger.Operator
var _ starter.Operator = &extractorDumajTabStarter{}

type extractorDumajTabStarter struct {
	interfaceKey joiner.InterfaceKey
}

func (edts *extractorDumajTabStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (edts *extractorDumajTabStarter) Prepare(cfg *config.Config, options common.Map) error {
	edts.interfaceKey = joiner.InterfaceKey(options.StringDefault("interface_key", string(extractor.InterfaceKey)))

	return nil
}

func (edts *extractorDumajTabStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	exytractorOp, err := New()
	if err != nil {
		return errors.Wrap(err, "can't init *extractorDumajTab{} as exchange.Operator")
	}

	if err = joinerOp.Join(exytractorOp, edts.interfaceKey); err != nil {
		return errors.Wrapf(err, "can't join *extractorDumajTab{} as exchange.Operator with key '%s'", edts.interfaceKey)
	}

	return nil
}
