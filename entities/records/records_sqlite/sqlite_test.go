package records_sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/common/common/apps"
	"github.com/pavlo67/common/common/db/db_sqlite"
	"github.com/pavlo67/common/common/starter"

	"github.com/pavlo67/data/entities/records"
)

func Testdb(t *testing.T) {
	_, cfgService, l := apps.PrepareTests(t, "../../../apps/", "test", "records_sqlite.log")
	require.NotNil(t, cfgService)

	components := []starter.Starter{
		{db_sqlite.Starter(), nil},
		{Starter(), nil},
	}

	joinerOp, err := starter.Run(components, cfgService, "CLI BUILD FOR TEST", l)
	require.NoError(t, err)
	require.NotNil(t, joinerOp)
	defer joinerOp.CloseAll()

	records.OperatorTestScenarioNoRBAC(t, joinerOp, l)
}
