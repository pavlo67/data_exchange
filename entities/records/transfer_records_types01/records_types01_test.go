package transfer_records_types01

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/apps"
	"github.com/pavlo67/common/common/db"
	"github.com/pavlo67/common/common/db/db_sqlite"
	"github.com/pavlo67/common/common/starter"

	"github.com/pavlo67/data/components/ns"
	"github.com/pavlo67/data/components/structures"
	"github.com/pavlo67/data/components/transfer"

	"github.com/pavlo67/data/entities/records"
	"github.com/pavlo67/data/entities/records/records_sqlite"

	"github.com/pavlo67/data/exchange/types01"
)

func TestTransferRecordsTypes01(t *testing.T) {
	err := os.Setenv("SHOW_CONNECTS", "1")
	require.NoError(t, err)

	_, cfgService, l := apps.PrepareTests(t, "../../../_environments/", "test", "")

	components := []starter.Starter{
		{db_sqlite.Starter(), nil},
		{records_sqlite.Starter(), nil},
		{Starter(), common.Map{"records_cleaner_key": records.InterfaceCleanerKey}},
	}

	label := "records_OPERATOR_PACK/TEST BUILD"
	joinerOp, err := starter.Run(components, cfgService, label, l)
	if err != nil {
		l.Fatal(err)
	}
	defer joinerOp.CloseAll()

	recordsOp, _ := joinerOp.Interface(records.InterfaceKey).(records.Operator)
	require.NotNil(t, recordsOp)

	recordsCleanerOp, _ := joinerOp.Interface(records.InterfaceCleanerKey).(db.Cleaner)
	require.NotNil(t, recordsCleanerOp)

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
		PackData: structures.NewDataAny([]types01.Record{
			{
				Content: types01.Content{
					Title:   "qwerwe",
					Summary: "rftyu",
					Type:    "wewe",
					Data:    "truyty",
				},
				Embedded: []types01.Content{
					{
						Title:   "qwerwe4",
						Summary: "rftyu444",
						Type:    "wew4444e",
						Data:    "truyty4444",
					},
				},
				ItemDescription: structures.ItemDescription{
					URN: "urn1",
					// History:   nil,
					// CreatedAt: time.Time{},
					// UpdatedAt: nil,
				},
			},
			{
				Content: types01.Content{
					Title:   "ert",
					Summary: "yuuuuuuuu",
					Type:    "eeee",
				},
				Embedded: []types01.Content{},
				ItemDescription: structures.ItemDescription{
					URN: "urn2",
					// History:   nil,
					// CreatedAt: time.Time{},
					// UpdatedAt: nil,
				},
			},
		}),
	}

	var params common.Map

	err = recordsCleanerOp.Clean(nil)
	require.NoError(t, err)

	// copyFinal, statFinal, dataFinal := transfer.TestOperator(t, transferOp, params, dataInitial, true, false)
	copyFinal, statFinal, dataFinal := transfer.TestOperator(t, transferOp, params, &dataInitial, true, false)

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
