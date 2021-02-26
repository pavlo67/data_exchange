package extractor_cli_settings

import (
	"fmt"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/joiner"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/starter"
	"github.com/pavlo67/data_exchange/components/transform"
	transform_records01_structure "github.com/pavlo67/data_exchange/components/transform/trabsform_records01_structure"
	"github.com/pavlo67/data_exchange/components/transform/transform_any_json"
	"github.com/pavlo67/data_exchange/components/transform/transform_structure_table"
	"github.com/pavlo67/data_exchange/components/transform/transform_table_csv"
)

func Components() ([]starter.Starter, error) {

	starters := []starter.Starter{
		// general purposes components
		//{control.Starter(), nil},
		//{connect_sqlite.Starter(), nil},

		// auth/persons components
		{transform_table_csv.Starter(), common.Map{"path": "../_extractor/_cliens.tab_"}},
		{transform_structure_table.Starter(), nil},
		{transform_records01_structure.Starter(), nil},
		{transform_any_json.Starter(), nil},
	}

	return starters, nil
}

type TransformNode struct {
	InterfaceKey joiner.InterfaceKey
	Operator     transform.Operator
}

func Run(joinerOp joiner.Operator, l logger.Operator) error {
	nodes := []TransformNode{
		{InterfaceKey: transform_table_csv.InterfaceKey},
		{InterfaceKey: transform_structure_table.InterfaceKey},
		{InterfaceKey: transform_records01_structure.InterfaceKey},
		{InterfaceKey: transform_any_json.InterfaceKey},
	}

	for i := range nodes {
		if nodes[i].Operator, _ = joinerOp.Interface(nodes[i].InterfaceKey).(transform.Operator); nodes[i].Operator == nil {
			return fmt.Errorf("no transform.Operator for key %s", nodes[i].InterfaceKey)
		}

	}

	var data interface{}
	var err error

	for i, node := range nodes {
		if i > 0 {
			if err = node.Operator.In(nil, data); err != nil {
				return fmt.Errorf("error on node.In %d (%s): %s", i, nodes[i].InterfaceKey, err)
			}
		}
		if data, err = node.Operator.Out(nil); err != nil {
			return fmt.Errorf("error on node.Out %d (%s): %s", i, nodes[i].InterfaceKey, err)
		}
	}

	l.Infof("DATA FINAL: %#v", data)

	return nil
}
