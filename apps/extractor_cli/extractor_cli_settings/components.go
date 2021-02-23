package extractor_cli_settings

import (
	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/config"
	"github.com/pavlo67/common/common/starter"
	"github.com/pavlo67/data_exchange/components/extractor/extractor_dumaj_tab"
)

func Components() ([]starter.Starter, error) {

	starters := []starter.Starter{
		// general purposes components
		//{control.Starter(), nil},
		//{connect_sqlite.Starter(), nil},

		// auth/persons components
		{extractor_dumaj_tab.Starter(), nil},

		// actions starter (connecting specific actions to the corresponding action managers)
		{Starter(), common.Map{
			"access":  config.Access{Path: "../_exchange/_cliens.tab_"},
			"path_to": "./",
		}},
	}

	return starters, nil
}
