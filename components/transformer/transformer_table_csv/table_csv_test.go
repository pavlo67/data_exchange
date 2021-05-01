package transformer_table_csv

import (
	"fmt"
	"testing"

	"github.com/pavlo67/data/components/structures"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/apps"
	"github.com/pavlo67/common/common/starter"

	"github.com/pavlo67/data/components/transformer"
)

func TestTransformTableCSV(t *testing.T) {
	_, cfgService, l := apps.PrepareTests(t, "../../../_environments/", "test", "")

	components := []starter.Starter{
		{Starter(), nil},
	}

	label := "TABLE_CSV/TEST BUILD"
	joinerOp, err := starter.Run(components, cfgService, label, l)
	if err != nil {
		l.Fatal(err)
	}
	defer joinerOp.CloseAll()

	transformOp, _ := joinerOp.Interface(InterfaceKey).(transformer.Operator)
	require.NotNil(t, transformOp)

	params := common.Map{
		"separator": "\t",
	}

	packInitial := structures.PackAny{
		PackDescription: structures.PackDescription{
			ItemDescription: structures.ItemDescription{
				URN: "test:test",
			},
		},
		PackData: structures.NewDataAny("as\tdfg r\tt/.jk\nrf\t .j;l'psa tproh\t\n\t\tnkcvbm/.sdgk'erlt;klghl\n;rkth;l"),
	}

	copyFinal, statFinal, dataFinal := transformer.TestOperator(t, transformOp, params, &packInitial, true, true)

	t.Logf("COPY (INTERNAL) FINAL: %#v", copyFinal)

	t.Logf("DATA (OUT) FINAL: %#v", dataFinal)

	if statFinalStringer, ok := statFinal.(fmt.Stringer); ok {
		t.Logf("STAT (INTERNAL) FINAL: %s", statFinalStringer.String())
	} else {
		t.Logf("STAT (INTERNAL) FINAL: %#v", statFinal)
	}

}
