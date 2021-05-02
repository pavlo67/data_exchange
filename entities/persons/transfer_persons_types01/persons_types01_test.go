package transfer_persons_types01

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/apps"
	"github.com/pavlo67/common/common/auth"
	"github.com/pavlo67/common/common/db"
	"github.com/pavlo67/common/common/db/db_sqlite"
	"github.com/pavlo67/common/common/rbac"
	"github.com/pavlo67/common/common/starter"

	"github.com/pavlo67/data/components/ns"
	"github.com/pavlo67/data/components/structures"
	"github.com/pavlo67/data/components/transfer"

	"github.com/pavlo67/data/entities/persons"
	"github.com/pavlo67/data/entities/persons/persons_sqlite"

	"github.com/pavlo67/data/exchange/types01"
)

func TestTransferPersonsTypes01(t *testing.T) {
	err := os.Setenv("SHOW_CONNECTS", "1")
	require.NoError(t, err)

	_, cfgService, l := apps.PrepareTests(t, "../../../_environments/", "test", "")

	components := []starter.Starter{
		{db_sqlite.Starter(), nil},
		{persons_sqlite.Starter(), nil},
		{Starter(), common.Map{"persons_cleaner_key": persons.InterfaceCleanerKey}},
	}

	label := "PERSONS_OPERATOR_PACK/TEST BUILD"
	joinerOp, err := starter.Run(components, cfgService, label, l)
	if err != nil {
		l.Fatal(err)
	}
	defer joinerOp.CloseAll()

	personsOp, _ := joinerOp.Interface(persons.InterfaceKey).(persons.Operator)
	require.NotNil(t, personsOp)

	personsCleanerOp, _ := joinerOp.Interface(persons.InterfaceCleanerKey).(db.Cleaner)
	require.NotNil(t, personsCleanerOp)

	transferOp, _ := joinerOp.Interface(InterfaceKey).(transfer.Operator)
	require.NotNil(t, transferOp)

	dataInitial := structures.PackAny{
		PackDescription: &structures.PackDescription{
			Fields: structures.Fields{},
			ItemDescription: structures.ItemDescription{
				URN: ns.URN("test:test" + strconv.FormatInt(time.Now().UnixNano(), 10)),
				// ErrorsMap: nil,
				// History:   nil,
				CreatedAt: time.Now(),
				// UpdatedAt: nil,
			},
		},
		PackData: structures.NewDataAny([]types01.Person{
			{
				Nickname: "wqerwqer",
				Roles:    nil,
				Creds:    auth.Creds{}, // auth.CredsEmail: "aaa@bbb.ccc"
				Info:     common.Map{"xxx": "yyy", "zzz": 777.},

				ItemDescription: structures.ItemDescription{
					// URN: "urn1",
					// History:   nil,
					// CreatedAt: time.Time{},
					// UpdatedAt: nil,
				},
			},
			{
				Nickname: "wqerwqer2",
				Roles:    rbac.Roles{rbac.RoleUser},
				Creds:    auth.Creds{}, // auth.CredsEmail: "aaa2@bbb.ccc"
				Info:     common.Map{"xxx2": "yyy", "zzz2": 222.},

				ItemDescription: structures.ItemDescription{
					// URN: "urn2",
					// History:   nil,
					// CreatedAt: time.Time{},
					// UpdatedAt: nil,
				},
			},
		}),
	}

	var params common.Map

	// copyFinal, statFinal, dataFinal := transfer.TestOperator(t, transferOp, params, dataInitial, true, false)
	copyFinal, statFinal, dataFinal := transfer.TestOperator(
		t, personsCleanerOp, transferOp, params, &dataInitial, true, false)

	//copyFinal, _ := transferOp.Copy(nil, params)
	t.Logf("COPY (INTERNAL) FINAL: %#v", copyFinal)

	//statFinal, _ := transferOp.Stat(nil, params)
	if statFinalStringer, ok := statFinal.(fmt.Stringer); ok {
		t.Logf("STAT (INTERNAL) FINAL: %s", statFinalStringer.String())
	} else {
		t.Logf("STAT (INTERNAL) FINAL: %#v", statFinal)
	}

	//dataFinal, _ := transferOp.Out(nil, params)
	t.Logf("DATA (OUT) FINAL: %#v", dataFinal)

}
