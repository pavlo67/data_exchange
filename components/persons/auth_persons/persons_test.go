package auth_persons

import (
	"testing"

	"github.com/pavlo67/common/common/apps"
	"github.com/pavlo67/common/common/auth"
	"github.com/pavlo67/common/common/starter"

	"github.com/pavlo67/data_exchange/components/persons/persons_fs"
)

func TestOperator(t *testing.T) {

	_, cfgService, l := apps.PrepareTests(
		t,
		"../../../apps/_environments/",
		"test",
		"", // "connect_test."+strconv.FormatInt(time.Now().Unix(), 10)+".log",
	)

	starters := []starter.Starter{
		{persons_fs.Starter(), nil},
		{Starter(), nil},
	}

	label := "CLI/TEST BUILD"
	joinerOp, err := starter.Run(starters, cfgService, label, l)
	if err != nil {
		l.Fatal(err)
	}
	defer joinerOp.CloseAll()

	authOp, _ := joinerOp.Interface(auth.InterfaceKey).(auth.Operator)
	if authOp == nil {
		l.Fatalf("no auth.Operator with key %s", auth.InterfaceKey)
	}

	auth.OperatorTestScenarioPassword(t, authOp)
}
