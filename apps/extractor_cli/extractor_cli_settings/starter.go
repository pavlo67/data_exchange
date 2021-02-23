package extractor_cli_settings

import (
	"fmt"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
	"github.com/pavlo67/data_exchange/components/extractor"
)

func Starter() starter.Operator {
	return &extractorStarter{}
}

var _ starter.Operator = &extractorStarter{}

type extractorStarter struct {
	access config.Access
	pathTo string
}

// --------------------------------------------------------------------------

var l logger.Operator

func (es *extractorStarter) Name() string {
	return logger.GetCallInfo().PackageName
}

func (es *extractorStarter) Prepare(cfg *config.Config, options common.Map) error {
	es.pathTo = options.StringDefault("path_to", "")
	if err := cfg.Value("extractor_cli", &es.access); err != nil {
		return err
	}

	return nil
}

func (es *extractorStarter) Run(joinerOp joiner.Operator) error {
	if l, _ = joinerOp.Interface(logger.InterfaceKey).(logger.Operator); l == nil {
		return fmt.Errorf("no logger.Operator with key %s", logger.InterfaceKey)
	}

	extractorOp, _ := joinerOp.Interface(extractor.InterfaceKey).(extractor.Operator)
	if extractorOp == nil {
		return fmt.Errorf("no extractor.Operator with key %s", extractor.InterfaceKey)
	}

	extractorOp.Draft(es.access, es.pathTo)

	return nil
}

// TODO!!! customize it
// if isHTTPS {
//	go http.ListenAndServe(":80", http.HandlerFunc(server_http.Redirect))
// }
